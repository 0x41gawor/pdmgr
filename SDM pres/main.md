# Plan ogólny

- Abstrakt (Takie tl;dr, słowa kluczowe)
- Geneza
- Problem
- Rozwiązanie
- Cel pracy
- Co do tej pory zrobiłem
- Co należy jeszcze zrobić

## Geneza
W 2022 rozpocząłem praktyki na stanowisku Inżyniera w Biurze Strategicznego Planowania i Standaryzacji Sieci. Jednym z zadań jakie wtedy dostałem było rozpoznanie kierunku rozwoju sieci mobilnych w kontekście 6G. Dla operatora ze strategicznego punktu widzenia takie zadanie ma sens, oraz wyniesiona wiedza może w jakiś sposób wpływać na długofalowe decyzje. 

Zapoznałem się wtedy z 4-oma whitepaper'ami. Od firm: Nokia, Huawei, Mediatek i Ericsson.

Bardzo duży nacisk w każdym z nich kładziony był na wykorzystanie sztucznej inteligencji w sieciach mobilnych. Technologie od lat się przeplatają, przyszedł i więc czas, aby wprowadzić sztuczną inteligencję do sieci. CHodzi tu o to, że w tym momencie sieci są utrzymywane przez ludzi. Każdy opeator posiada coś takiego jak Network Operations Center (NOC). Głównym jego celem jest zapewnienie ciągłości działania sieci oraz szybkie regaowanie na awarie lub inne problemy. Głównie tutaj sztuczna inteligencja mogłaby zastąpić człowieka, ale również coś takiego jak:
- analityka sieci (wyciąganie wniosków na jej postawie),
- automatyczne wprowadzanie w niej zmian (uruchamianie nowych usług szytych na potrzeby klientów),
- zarządzanie zasobami sieci (wyłączanie niedociążonych urządzeń (żeby być eko), alokowanie zasobów chmurowych w odpowiednich geolokalizacjach itp.)

Tu trzeba dodać, że sieci mobilne jako niesamowicie złożone twory, obejmują wiele dziedzin naukowych, bedące czymś bardzo skomplikowanych są rzetelnie standaryzowane. Organizacją standaryzująca jest 3GPP, która zrzesza pracowników wielu firm (m.in. tych wymienionych wcześniej od whitepapers). 3GPP jest łączy 7 ciał standaryzacyjnych z całego świata: USA, Japonia, Chiny, Korea, Indie oraz Europa. Europejskim reprezentantem jest ETSI - The European Telecommunications Standards Institute (ETSI).

ETSI z kolei podzielone jest na komitety. A jednym z nich jest ENI - Experiential Networked Intelligence, które właśnie zastanawia się i bada jak ulepszyć zycie operatora sieci mobilnej poprzez użycie sztucznej inteligencji w formie zamkniętych pętli sterowania bazujących na świadomych konktekstu i opartych na metadanych politykach (ang. "policy').

### dygresja czym jest pętla sterowania


https://www.etsi.org/technologies/experiential-networked-intelligence

No i właśnie ENI przedstawia dokument omawiający znane już ludzkości z innych dziedzin pętle sterowania. 

https://www.etsi.org/deliver/etsi_gr/ENI/001_099/017/02.01.01_60/gr_ENI017v020101p.pdf

Tak, aby na ich podstawie wypracować architekturę systemu odpowiadającego na postawione zadanie.

## Problem

Każde inżynierskie zadanie wykonane w historia ludzkości można przedstawić w modelu problem-rozwiązanie. Zacznijmy do zdefiniowania problemu, który praca magisterska chce rozwiązać. Otóż ENI definiuje jedynie (jak to standard architekturę), jest ona wysokokopoziomowa i abstrakcyjna. Musi być ogólna, aby znalzała zastosowanie. Standardy zazwyczaj nie mówią nic o implementacji, one tylko specyfikują zachowanie oraz interfejsy. Jakie jest następny krok? Co jest problemem? Przybliżenie tej abstrakcji na implementację.

Spójrzmy na architekturę ENI. Jak teraz zaimplementować bloczek "ENI System"?

Potrzebujemy sposobu, ogólnej struktury (ang. "framework") na zamodelowanie dowolnej zamkniętej pętli sterowania. Co więcej postawione są pewne wymagania na tę pętlę, takie jak:
- świadomość kontekstu
- zarządzanie wiedzą
- procesowanie kognitywne
- bazująca na modelu
- zarządzana polisami

Te wszystkie ciężkie terminy są zdefiniowane w standardach ENI. My jedynie musimy umożliwić, aby pętle zaimplementowane według naszego frameworku mogły takie być.

> Uważny słuchacz moze zapytać, ok ale gdzie to AI?

## Rozwiązanie

Impementacja programistyczna aplikacji będącej pętlą sterowania nie jest w żaden sposób ogólna, jako że jej logika oraz polisy zakodowane by zostały w kodzie aplikacji. Takie rozwiązanie odpada na starcie i służy jedynie jako obrazowy punkt wyjścia. 

Należy więc zaprogramować platformę, która jest ogólna. W kodzie zakodowane są jedynie ogólne wzorce logiki pętli oraz interpter polis. Platforma ta pozwalałaby na zdefiniowanie w niej i wykonanie dowolnej z pętli sterowania oraz na przekazanie jej dowolnych polis/polityk. Dodatkowo posiadała by jasno zdefiniowane interfejsy, tak aby można było podłączyć do niej moduły AI, które realizowałyby: świadomość kontekstu, zarządzanie wiedzą, procesowanie kognitywne itp. Również musi ona pozwalać na zdefiniowanie modelu na jakim pracujemy. Dodatkowo wszystkie interfejsy i standardy z jakich by korzystała muszą dobrze znane szeroko pojętej społeczności, która by z niej korzystała. 

Należy tu też pamiętać, że użytkownikiem takiej platformy jest operator sieci mobilnej. Inżynier do spraw utrzymania sieci, lub twórca na niej usług biznesowych. Nie jest to z żadnym wypadku programista. Platforma więc musi być zrozumiała i łatwa do korzystania dla takiej osoby. 

No dobrze w takim razie czas zabrać się do kodowania .... Ale hola hola. Może zamiast budować wielką platformę od zera, warto zastanowić się czy nie można użyć istniejącego już narzędzia lub zbioru narzędzi. 

I tutaj potocznie mówiąc cały na biało przychodzi Kubernetes. Kubernetes u swoich podstaw działa w oparciu o zamknięte pętle sterowania, dodatkowo jest to platforma doskanale już znana i wspierana w społeczności sieci mobilnych. 

https://kubernetes.io/docs/concepts/architecture/controller/

https://kubernetes.io/docs/concepts/extend-kubernetes/operator/

https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/

Każdy API object taki jak Pod, Deployment, Service, Replicaset posiada swoją zamkniętą pętle sterowania opartą o proces sterownika (ang. "controller"), który znajduje się w podzie "Controller Manager" w płaszczyźnie sterowania. Kubernetes posiada mechanizmy rozszerzalności i tak .. Mechanizm Custom Resource Definitions (CRD) pozwala defniowanować swoje własne API Objects a następnie tworzyć ich instancję. Każdy API Object musi mieć swój sterownik i w tym celu Kubernetes udostępnie Operator Pattern czyli mechanizm developowania swojego własnego sterownika dla swoich własnych API Objects. 

Taki był początkowy wybór reszta narzędzi wychodzi w trakcie pracowni magisterskiej i również zostanie omówiona.

## Cel pracy

Reasumując więc celem pracy jest:

Zaproponowanie framework'u, sposobu, ogólnego schematu definiowania zamkniętych pętli sterowania w oparciu o platformę Kubernetes, a w szczególności o jej mechanizmy Custom Resource Definitions oraz Operatorów. 

Wymagania na platformę/framework:
- pozwala zdefiniować logikę dowolnej zamkniętej pętli sterowania
- pozwala na komunikacje z zewnętrznymi komponentami (będą to bloki obliczniowe AI, celem jest delegacja części logiki (tej nie będącej generyczną))
- użytkownikiem jest osoba z jedynie pobiężna wiedzą techniczną, musi ona jedynie nauczyć się jak za pomocą platformy wyrażać potrzebną jej logikę biznesową, nie jest to w żandym wypadku programista
- pozwala na rozszerzalność, tam gdzie kod operatora ogranicza użytkownika i jest uzasadnienie biznesowe, możliwe jest dopisane "snippetu" kodu przez zespół operatora sieci mobilnej

## Co do tej pory zrobiłem

[raport](raport.pdf)

- Zapoznanie się z koncepcją zamkniętej pętli sterowania wg. ETSI ENI
- Zapoznanie się z platformą Kuberntes (solidne podstawy potrzebne do zrozumienia CRDs i Operator Pattern)
- Utworzenie środowiska deweloperskiego
- Praktyczne zapoznanie się z CRDs
- Zapoznanie się ze środowiskiem Kubebuilder
- Praktyczne zapoznanie się z Operator Pattern poprzez PoC: Operator komunikujący się z zewnętrznym komponentem poprzez REST
- Zapoznanie się z gRPC jako alternatywny dla HTTP sposób komunikacji z komponentem zewnętrznym
- Zapoznanie się z Open Policy Agent jako narzędzie do budowania polis/poltyk
- Koceptualny podział pętli na bloki funkcjonalne (na pierwowzorze pętli ENI)
- Wyobrażenie sobie jak zaimplementować bloki funkcjonalne pętli w Kubernetes (https://github.com/0x41gawor/cloopdemo1/blob/master/main.md)
  - Następstwem tego kroku jest wypracowanie architektury platformy
- Wymyślenie przykładu, na którym pracować będe podczas fazy developmentu platformy
- Rozpoczęcie fazy developmentu (Tu już jest implementacja kodu operatorów etc.)

## Rozbijemy teraz każdy z kroków 

## Co należy jeszcze zrobić

- Zamknięcie fazy developmentu (Implementacja kodu operatorów), a w tym:
  - Wbudowanie odwołań do OPA w operatorze
  - Oóglne wbudowanie odwołań do komponentów zewnętrznych w operatorze
  - Wymyślenie dobrej konwencji nazewnictwa, tak aby z nazwy można było łatwo/automatycznie wywieść nazwy Custom Resouce, Operatorów, URL w OPA itp.
  - Zdefiniowanie granic. Co jest zakresem kodu operatora a co polityką w OPA
- Próba dojścia do generycznych operatorów, tak aby jedynie plikami CR, politykami OPA oraz komponentami zewnętrznymi definiować logikę pętli. Dojście do stanu "data-driven".
- Przeprowadzania Proof Of Concept użycia naszej platformy (dobre by było kilka przykładów)
- Stworzenie odpowiedniej dokumentacji platformy, która stanowiłaby jako instrukcji użytkowania
