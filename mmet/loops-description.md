# Loops description

This document presents more technical (algorithms, data structures) description about loops used in mmet.

Please review [business-case.md](business-case.md) before proceeding.

Status: Doc is still under construction.

## Architecture 
Let's first discuss the components of our system before delving into the algorithms used in the loops.

We can distinguish 4 components:
- `mmet` - Represents the MME Network. From the perspective of the closed control loop, it acts as the **managed system**.
- `translation agent` - Serves as a bridge between the loops and the external world; in our case, the external world is `mmet`. It:
  - Retrieves data from `mmet`
  - Sends action commands to `mmet`
- `reactive loop` - This loop ensures and enforces the demanded traffic distribution.
- `deliberative loop` - Supervising the reactive loop, this component can modify distribution values.



Table below give some insights about technical implementation/placement of these components.

| Component name:   | Implemented as:                                        | Belongs to our platform? |
| ----------------- | ------------------------------------------------------ | ------------------------ |
| mmet              | Linux process, likely a Go application                 | No (external)            |
| translation agent | Kubernetes pod, likely a Go application in a container | Yes (entry point)        |
| reactive loop     | Custom Resource API Objects in Kubernetes              | Yes                      |
| deliberative loop | Custom Resource API Objects in Kubernetes              | Yes                      |

## mmet

### General descr

This Linux process generates a quadruple each round, where each value represents the number of sessions currently served by the corresponding MME node.

For example, it produces JSON as below:
```json
{
  "Gdansk": 10,
  "Poznan": 12,
  "Warsaw": 18,
  "Krakow": 6
}
```

Each round, a random number (which can be negative) is added to each node's count. 

This results in a trajectory like the following

```sh
```sh
# roundNumber. Gdansk-Poznan-Warsaw-Cracow
1. 10-12-18-6 (1, 0, -2, 2)
2. 11-12-16-8 (-1, -2, 1, 1)
3. 10-10-17-9 (-2,-2,3,-3)
4. 8-8-20-6 (-2,-1,-2,2)
5. 6-7-18-8
```

Additionally, the program can receive commands specifying traffic movement from one node to another ("move `x` units of traffic from node `a` to node `b`"), such as:

```json
{
  "count": 2,
  "from": "Warsaw",
  "to": "Poznan"
}
```

By default, rounds occur every minute.

Data generated each round are sent via HTTP to a pre-configured endpoint in the JSON format shown above. Each round begins with a log statement:

```sh
"Round 1: {Gdansk: 10, Poznan: 12, Warsaw: 18, Krakow: 6}"
```

Furthermore, the program listens on HTTP port 4545 (/api/move) and logs incoming move commands:
```sh
"Got move command: {count: 2, from: Warsaw, to: Poznan}"
```

The initial configuration of distribution values is hardcoded in an accessible location.

### Implementation

See [mmet.go](mmet.go).

## Translation Agent

### General descr

Translation Agent is a linux process that:

- Operates an HTTP server and listens to requests from `mmet`
- Utilizes an HTTP client to send Move Commands to `mmet`
- Utilizes the kube-api-server client to push data into the `reactive loop`

It is containerized using Docker and deployed as a pod in a Kubernetes cluster (Minikube).

