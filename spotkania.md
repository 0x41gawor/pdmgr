# 24 03 27

### Logika pętli deklaratywnie, zero kodu GO

Logika pętli deklaratywnie - chcemy dać framework taki, który będzie mógł używać ktoś komu wytłumaczymy tylko jak używać naszych Custom Resource. Nie będzie trzeba pisać żadnego kodu w Go. Jakiś półtechniczny Service Manager będzie tylko wypełniał CRki. To będzie jego interfejs do naszych pętli. Ten ktoś będzie w klastrze zwykłym userem, który może stawiać instancje swoich obiektów. I właśnie on będzie stawiał takie instancje, żeby zadziały się nasze pętle. Jedna instancja naszych CR jest nastawiona na coś konkretnego, na tym poziomie już jest specjalizacja. Admin naszych pętli jest tylko userem w klastrze. Nie chcemy, żeby pisał jakikolwiek kod w GO. 

### User naszego frameworka jest półtechniczny

My dajemy mu zbiór zasad i template'ty i mówimy jak ma postępować, aby zrobić pętle taką a taką. 
Userem jest jakiś Service Manager, półtechniczny człowiek, który w klastrze ma uprawnienia do tworzenia API Objects. 

On w głowie ma logikę biznesową pętli, a my dajemy mu sposób jak może ją przenieść na rzeczywisty świat.

Nie jest to programista, nie chcemy by pisał jakikolwiek kod w GO.

### Extensibility

Generalnie nie chcemy, żeby w naszym frameworku/platformie ktoś musiał pisać kod, ale no jakąś rozszerzalność należy dać. Jest to dobra praktyka. I tak jak K8s daje ją właśnie przed CRD i Operatory, to my damy możliwość dopisania snippetu kodu do operatora.

