# Loops description

This document presents more technical (algorithms, data structures) description about loops used in mmet.

Please get familiar with [business-case.md](business-case.md) first.

## Architecture 
Before diving into explanation of algorithms used in loops let's talk about the components of our system.

We can distinguish 4 components:
- `mmet` - MME Network. From closed control loop's perspective it stands as **managed system**.
- `translation agent` - in general, bridge between loops and external world, in our case external world is mmet. It:
  - pulls/listens for data from mmet
  - pushes action commands to mmet
  
- `reactive loop` - this loop will ensure/enforce demanded traffic distribution. 
- `deliberative loop` - this loop, while supervising the reactive one, can change distribution values. 



Table below give some insights about technical implementation/placement of these components.

| Component name:   | Implemented as:                                          | Belongs to our platform? |
| ----------------- | -------------------------------------------------------- | ------------------------ |
| mmet              | Linux process, probably a go application                 | No, external world       |
| translation agent | Kubernetes pod, probably a go application in a container | Yes, entry point         |
| reactive loop     | Custom Resource API Objects in Kubernetes                | Yes                      |
| deliberative loop | Custom Resource API Objects in Kubernetes                | Yes                      |

## mmet

### General descr

This linux process gives some quadruple each round. Each value in quadruple represents a number of session served at the moment (at this round) by corresponding MME node.

E.g. it will produce json as below:
```json
{
  "Gdansk": 10,
  "Poznan": 12,
  "Warsaw": 18,
  "Krakow": 6
}
```

At each round a number (can be less than 0) that will be added to each node count is randomly chosen. 

It can result in such trait:
```sh
```sh
# roundNumber. Gdansk-Poznan-Warsaw-Cracow
1. 10-12-18-6 (1, 0, -2, 2)
2. 11-12-16-8 (-1, -2, 1, 1)
3. 10-10-17-9 (-2,-2,3,-3)
4. 8-8-20-6 (-2,-1,-2,2)
5. 6-7-18-8
```

Beside form generating values a program also can receive some input, that says "move `x` units of traffic from node `a` to node `b`".

E.g.
```json
{
  "count": 2,
  "from": "Warsaw",
  "to": "Poznan"
}
```

By default let's assume that round happens each minute. 

The values generated for each round are send via http to pre-configured endpoint in a json format presented above. 
Each round start with Log:
```sh
"Round 1: {Gdansk: 10, Poznan: 12, Warsaw: 18, Krakow: 6}"
```

Additionally, program listens on http port 4545 (api/move) and when incoming message move comes it logs:
```sh
"Got move command: {count: 2, from: Warsaw, to: Poznan}"
```

The starting configuration of distribution values is hardcoded in some accessible place.