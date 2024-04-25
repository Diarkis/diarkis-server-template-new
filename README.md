# Overview

Diarkis server cluster is made up with HTTP, TCP, UDP and WebSocket servers.

Each protocol servers run independently within the cluster, but you do not have to have all protocols.

Only HTTP server is required in the cluster and the rest of the servers should be chosen according to your application's requirements.

# Structure

```
───┬─ servers/ ─┬──── http/main.go      [HTTP server main]
   │            │
   │            ├──── udp/main.go       [UDP server main]
   │            │
   │            ├──── tcp/main.go       [TCP server main]
   │            │
   │            └──── connector/main.go [Connector server main]
   │
   ├─ bot/ [Bot clients for stress test]
   │
   ├─ mars/ ───────── main.go
   │
   │
   ├─ healthcheck/ ── main.go
   │
   │
   ├─ configs/ ─┬──── shared/ [Shared configuration directory] ────────────────────┬─ field.json
   │            │                                                                  ├─ group.json
   │            ├──── http/     [HTTP configuration directory] ──────── main.json  ├─ log.json
   │            │                                                                  ├─ matching.json
   │            ├──── udp/      [UDP configuration directory]  ──────── main.json  └─ mesh.json
   │            │
   │            ├──── tcp/      [TCP configuration directory]  ──────── main.json
   │            │
   │            └──── connector [Connector configuration directory] ─── main.json
   │
   ├─ cmds/  [Custom client command directory] ────────────────┬── main.go [Entry point for all cmds]
   │                                                           │
   │                                                           ├── http   ──────────────────────────────────────┬─── main.go
   ├─ lib/   [Shared library directory]                        ├── room   ──────────────────────────── main.go  └─── matching.go
   │                                                           ├── group  ──────────────────────────── main.go
   ├─ bin/   [Built server binary directory]                   ├── field  ──────────────────────────── main.go
   │                                                           └── custom ──────────────────────────── main.go
   ├ cloud/ [Cloud Infra Manual for Diarkis]
   │   
   ├ k8s/ [Contains k8s manifest for Diarkis]
   │   
   ├  puffer/ [Contains package definition generator]
   │
   ├ build.yml [Build configuration file for diarkis-cli]
   │
   └─ go.mod [Go module file for the project]
```

# Building Server Binaries

Use the following commands to build the server binaries:

## Build Servers for Local Use

The following command will be building the servers for your local machine.

```
make local-build
```

### local-build.yml

This file controls which server binaries to build, which Go version to use, and which architecture and OS to build the servers for etc.

### Start a Server locally

The following command will start the binary on your local machine.

```
make server target=$(udp|tcp|http|mars)
```

**NOTE** You must have MARS and HTTP server running in order to create Diarkis server cluster properly.

# Server Entry Points

## HTTP server

```
servers/http/main.go
```

## UDP server

```
servers/udp/main.go
```

## TCP server

```
servers/tcp/main.go
```

## Connector server

```
servers/connector/main.go
```

# MARS

Diarkis server cluster requires it's unique server called MARS server.

You need to simply build MARS server and deploy it along with other Diarkis servers.

## Starting MARS Server With Configuration File

By default, there is MARS configuration JSON file: `configs/mars/main.json`

```
./mars <path to the config JSON file>
```

# Health Check

Diarkis server needs to have health check. This template provides the source to build the health check binary.

The build will be automatically executed when you execute our make commands.

# Stress Test Bots

Diarkis server template comes with simple bots for stress tests.

All bots can be built using diarkis-cli

## MatchMaker Bot

This bot uses MatchMaker built-in commands to perform search and add.

```
bots/matchmaker/main.go
```

### How To Use MatchMaker Bot

```
./remote_bin/bot-matchmaker {HTTP endpoint:port} {How many bots to spawn} {Raito of hosts from 0% to 100%} {Search interval in milliseconds}
```

Example:

```
# Bot executable binary     Target host    Number of bots  Host ratio (30%)  Search interval (500ms)
./remote_bin/bot-matchmaker 127.0.0.1:7000 1000            30                500
```

# Commands

This is where you add your custom commands.

## UDP and TCP

```
cmds/
```

## WebSocket

```
ws_cmds/
```

We recommend that packets be defined and implemented using `puffer`.
We store puffer in the puffer directory and provide usage and examples.

# Configurations

This is where you place your configuration JSON files.

```
configs/
```

# MatchMaker With UDP and/or TCP Server

There are two sample matchmaking implemented for UDP and/or TCP server.

```
cmds/custom/main.go
```

## Sample Matching IDs

- RankMatch

  - `rank` is range of 5

- RateAndPlay

  - `rate` is range of 1

  - `play` is range of 1

## MatchMaker Add Command For UDP/RUDP and TCP

MatchMaker add command creates a new room and adds that to MatchMaker to be searched and found.

Client receives a response with `ver:2` and `cmd:100` to evaluate success or failure of the command.

Client raises `On Room Creation` event if successful.

### Command Version and ID

```
ver: 2
cmd: 100
```

### Payload

Endianess is `Big Endian`.

```
+--------+-----------+-----------------+--------------------+
|   TTL  |   *Size   |   String List   |    Property Map    |
+--------+-----------+-----------------+--------------------+
| 8 byte |   2 byte  |      *Size      |   variable size    |
+--------+-----------+-----------------+--------------------+
| 0    7 | 8       9 | 10   10 + *size | 10 + *size + 1 ... |
+--------+-----------+-----------------+--------------------+
```

**String List**
```
+---------+-------------+----------+-----------+
|  *size  | Matching ID |  **size  | Unique ID |
+---------+-------------+----------+-----------+
|  4 byte |    *size    |  4 byte  |  **size   |
+---------+-------------+----------+-----------+
```

**Property Map**

Property Value, size, and Property Name may repeat as a set of data.

```
+----------------+--------+---------------+
| Property Value | *size  | Property Name |
+----------------+--------+---------------+
|     4 byte     | 2 byte |     *size     |
+----------------+--------+---------------+
| 0            3 | 4    5 | 6   6 + *size |
+----------------+--------+---------------+
```

## MatchMaker Search Command For UDP/RUDP and TCP

MatchMaker search finds rooms that matches the given properties (conditions) and join the found room.

Client receives a response with `ver:2` and `cmd:102` to evaluate success or failure of the command.

Remote clients that already matched receives a server push with `ver:2` and `cmd:103` to notify the matched room is now full.

Remote clients that already matched raise `On Member Join` event on successful search.

### Command Version and ID

```
ver: 2
cmd: 102
```

### Payload

```
+--------+-----------------+---------------+
| *size  |   String List   |  Propery Map  |
+--------+-----------------+---------------+
| 2 byte |      *size      |    Variable   |
+--------+-----------------+---------------+
| 0    1 | 2     2 + *size | 2 + *size ... |
+--------+-----------------+---------------+
```

**String List**

size and Matching ID may repeat as a set of data.

```
+---------+-------------+
|  *size  | Matching ID |
+---------+-------------+
|  4 byte |    *size    |
+---------+-------------+
```

**Property Map**

Property Value, size, and Property Name may repeat as a data set.

```
+----------------+--------+---------------+
| Property Value | *size  | Property Name |
+----------------+--------+---------------+
|     4 byte     | 2 byte |     *size     |
+----------------+--------+---------------+
| 0            3 | 4    5 | 6   6 + *size |
+----------------+--------+---------------+
```

## P2P Address Report Command For UDP/RUDP and TCP

The command is used to report the client's public address meant for peer-to-peer communication.

This command assumes the client has joined a room.

### Command Version and ID

```
ver: 2
cmd: 110
```

### Payload

Client Address is a byte array encoded UTF8 string.

```
+---------------+----------------+
| Custom Header | Client Address |
+---------------+----------------+
|     5 byte    |    Variable    |
+---------------+----------------+
| 0           4 | 5          ... |
+---------------+----------------+
```

## P2P Initialize Command For UDP/RUDP and TCP

Starts peer-to-peer communication with all members of the room the client is in.

Client receives a response with `ver:2` and `cmd:111` to evaluate success or failure of the command.

All remote clients that are members of the room raise `On Member Broadcast` with a list of other client's addresses.

The clients may use those addresses to initiate peer-to-peer communication immediately.

### Command Version and ID

```
ver: 2
cmd: 111
```

### Payload

Empty payload.

### Payload Of On Member Broadcast

**String List**

size and Client Address may repeat as a set of data.

```
+---------+----------------+
|  *size  | Client Address |
+---------+----------------+
|  4 byte |      *size     |
+---------+----------------+
```

---

## Output Client Error Log On UDP/RUDP and TCP Server

This custom command expects the client to send error log and output the error log as an error log on the server.

### Command Version and ID

```
ver: 2
cmd: 12
```

### Payload

The payload should be a UTF8 encoded string.

---

# Creating A New Room Via Diarkis HTTP Server

You may create an empty room from Diarkis HTTP server.

```
POST /room/create/:serverType/:maxMembers/:ttl/:interval
```

## Parameters

- `serverType` is to choose which server to create a new room in. Valid types are: `udp`, `tcp`, and `ws`.

- `maxMembers` is a maximum client members allowed in the new room.

- `ttl` is a TTL value for the empty room to be kept. TTL is in seconds.

- `interval` is an interval of room broadcast in milliseconds. The room will buffer broadcast message on every interval.

---

# MatchMaker With HTTP Server

This is where you define your own MatchMaker rules.

The template provides HTTP API endpoints, but you may implement UDP, TCP, WebSocket commands for MatchMaker as well.

```
cmds/http/matching.go
```

## Default HTTP API Endpoints

We have pre-define rules of `rank` and `score`.

`rank` and `score` are also the `mmID` of each rule.

### Adding data to MatchMaker pool to be searched

```
POST /mm/add/:mmID/:uniqueID/:ttl
```

- `mmID` is the ID of MatchMaker rule.

- `uniqueID` is the unique ID of the data that is to be added to MatchMaker pool.

- `ttl` is the TTL of the data that is to be added to MatchMaker pool.

#### Request Body

- `props` is the JSON data to represent conditions of the data. This is used by search.

- `metadata` is the JSON data to be returned when search returns with results.

### Searching candidates from MatchMaker pool

```
POST /mm/search/:mmIDs/:limit
```

- `mmIDs` is a comma separated MatchMaker rule IDs.

- `limit` defines how many matching results you expect.

#### Request Body

- `props` is the JSON data representation of search conditions.

# Custom Commands

This is where you implement your own custom commands for TCP, UDP/RUDP.

```
/cmds/custom/main.go
```

This is where you implement your own custom commands for WebSocket.

```
ws_cmds/custom/main.go
```
