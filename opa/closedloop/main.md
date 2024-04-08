# Business case

Generalnie chodzi o to, że operator oddelegowuje swoją logikę control loop do jakiegoś zewnętrznego komponentu. Tym zewnętrznym tutaj komponentem jest OPA.

Business case w tym przykładzie jest taki, że operator wysyła stały input do OPA. Operator monitoruje po prostu jakiś resource (wiesz, że to jest tak, że reconcillation loop odpala się, gdy coś się zmieni - taki jest trigger events).

```json
{
    "cpu": <percentage value of utilization>,
    "memory": <percentage value of utilization>
}
```

No i OPA po pierwsze sprawdza czy thresholdy na daną metrykę zostały przekroczone czy nie. Zwraca, więc info o tym, ale przede wszystkim zwraca decyzję czy należy coś z jakąś metryką robić. Ten prosty przykład ma dwie metryki, więc decyzja może przyjąć jedną z wartości z tej trójki:

```go
["none", "cpu", "memory"]
```

Operator dostając taką odpowiedź będzie wiedział już co robić. Logika jak się ratować gdy cpu/memory zostanie przekroczone już będzie zapisana w nim w golangu.

Tak, więc output zwraca

```json
{
    "decision": <action name to perform or none>,
    "monitoring": {
  		"cpu": <boolean value if threshhold exceeded or not>
    	"memory": <boolean value if threshhold exceeded or not>
	}
}
```

But ofc. this is only exemplary set of metrics to monitor. The rego rules should be prepared for any number of metrics. Co jeśli jakiś operator obserwuje dużo więcej rzeczy.

# data.json

```json
{
    "Decisionpolicies": {
        "Decisiontype": "Priority",
        "Priorityspec": {
            "Priorityrank": {
                "rank-1": "cpu",
                "rank-2": "memory"
            },
            "Prioritytype": "Basic",
            "Time": "2023-12-01 21:51:58.427048"
        }
    },
    "Monitoringpolicies": {
        "Data": {
            "MonitoringData-1": "memory",
            "MonitoringData-2": "cpu"
        },
        "Time": "2023-12-01 21:51:58.427048",
        "Tresholdkind": {
            "MonitoringData-1-thresholdkind": "inferior",
            "MonitoringData-2-thresholdkind": "inferior"
        },
        "Tresholdvalue": {
            "MonitoringData-1-thresholdvalue": 50,
            "MonitoringData-2-thresholdvalue": 5
        }
    }
}
```

data dzielimy na dwie sekcje, jedna dotyczy tego jaką podejmiemy decyzję druga jak z tym monitorowaniem. Dla każdego operator-case trzeba będzie wziąć taki template'owy plik i wypełnić jakie metryki są obserwowane przezeń. Modyfikowane bedą pola json:

- `Decisionpolicies.PrioritySpec.PriorityRank`
- Oraz Cały `Monitoringpolicies` prócz Time. Tam dla każdej metryki należy ustawić:
  - Jej nazwę
  - Czy ma być poniżej/równa/czy powyżej zadanej wartości threshold
  - Zadany threshold.

> Ja bym tu podmienił frazę `MonitoringData` na `Metric`.

# rego rules

Plik rego dzieli się na 3 części:

- ustawienie zmiennej bool `cpu`
- ustawienie zmiennej bool memory
- wyprodukowanie na ich podstawie outputu

## ustawienie zmiennej bool `cpu`

```rego
default cpu := false

# są trzy przypadki thresholda metryki, albo ma być ona poniżej, powyżej albo równa
# jeśli jedno z tych 3 będzie true to cpu będzie true
cpu if {
    some cpu_idx									
    data.Monitoringpolicies.Data[cpu_idx] == "cpu"					
    cpu_idx_tresh := sprintf("%v%v", [cpu_idx, "-thresholdvalue"]) #sprintf to funkcja do formatowania stringów
    cpu_idx_kind := sprintf("%v%v", [cpu_idx, "-thresholdkind"])   # tu użyta do konkatenacji 
    data.Monitoringpolicies.Tresholdkind[cpu_idx_kind] == "inferior"
    input.cpu > data.Monitoringpolicies.Tresholdvalue[cpu_idx_tresh]
}

cpu if {
    some cpu_idx										# if for some cpu_idx - jeśli dla jakiegoś indeksu
    data.Monitoringpolicies.Data[cpu_idx] == "cpu"		#  w data.Monitoringpolicies.Data znajdziemy takie pole, że =="cpu"
    cpu_idx_tresh := sprintf("%v%v", [cpu_idx, "-thresholdvalue"])
    cpu_idx_kind := sprintf("%v%v", [cpu_idx, "-thresholdkind"])
    data.Monitoringpolicies.Tresholdkind[cpu_idx_kind] == "superior" # i na dodatek tu jest tak
    input.cpu < data.Monitoringpolicies.Tresholdvalue[cpu_idx_tresh] # oraz jeszcze o to, to wtedy cały rule cpu przyjmie true
}

cpu if {
    some cpu_idx
    data.Monitoringpolicies.Data[cpu_idx] == "cpu"
    cpu_idx_tresh := sprintf("%v%v", [cpu_idx, "-thresholdvalue"])
    cpu_idx_kind := sprintf("%v%v", [cpu_idx, "-thresholdkind"])
    data.Monitoringpolicies.Tresholdkind[cpu_idx_kind] == "uniform"
    input.cpu == data.Monitoringpolicies.Tresholdvalue[cpu_idx_tresh]
}
```

## ustawienie zmiennej bool `memory`

```rego
default memory := false

memory if {
    some memory_idx
    data.Monitoringpolicies.Data[memory_idx] == "memory"
    memory_idx_tresh := sprintf("%v%v", [memory_idx, "-thresholdvalue"])
    memory_idx_kind := sprintf("%v%v", [memory_idx, "-thresholdkind"])
    data.Monitoringpolicies.Tresholdkind[memory_idx_kind] == "inferior"
    input.memory > data.Monitoringpolicies.Tresholdvalue[memory_idx_tresh]
}

memory if {
    some memory_idx
    data.Monitoringpolicies.Data[memory_idx] == "memory"
    memory_idx_tresh := sprintf("%v%v", [memory_idx, "-thresholdvalue"])
    memory_idx_kind := sprintf("%v%v", [memory_idx, "-thresholdkind"])
    data.Monitoringpolicies.Tresholdkind[memory_idx_kind] == "superior"
    input.memory < data.Monitoringpolicies.Tresholdvalue[memory_idx_tresh]
}

memory if {
    some memory_idx
    data.Monitoringpolicies.Data[memory_idx] == "memory"
    memory_idx_tresh := sprintf("%v%v", [memory_idx, "-thresholdvalue"])
    memory_idx_kind := sprintf("%v%v", [memory_idx, "-thresholdkind"])
    data.Monitoringpolicies.Tresholdkind[memory_idx_kind] == "uniform"
    input.memory == data.Monitoringpolicies.Tresholdvalue[memory_idx_tresh]
}
```

Same as above.

## Wyprodukowanie output

```rego
monitoring := {"cpu": cpu, "memory": memory}
default decision := "none"

decision := result if {
    data.Decisionpolicies.Decisiontype == "Priority"
    monitoring.cpu == true
    monitoring.memory == true
    result := data.Decisionpolicies.Priorityspec.Priorityrank["rank-1"]
}

decision := result if {
    data.Decisionpolicies.Decisiontype == "Priority"
    monitoring.cpu == true
    monitoring.memory == false
    result := "cpu"
}

decision := result if {
    data.Decisionpolicies.Decisiontype == "Priority"
    monitoring.cpu == false
    monitoring.memory == true
    result := "memory"
}
```

//TODO poznaj te składnie