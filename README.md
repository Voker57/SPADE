# SPADE
**SPADE**: Digging into **S**elective and **PA**rtial **DE**cryption using Functional Encryption


## Limitations
Curators should specify the number of users and the maximum number of entries for each user.
Users will encrypt their data as a vector of integers that cannot have **zero** values.

**Note:** Be careful when defining users' data vector using `make()`; 
it assigns a **zero** to the elements when initializing them.

## Changing the protobuf structure
If you want to modify the protobuf structure first change the following file:

    ./spadeproto/spade.proto

and then run the proto compiler command to generate the new protobuf files:

    protoc --go_out=. --go-grpc_out=. spade.proto 

    
## Testing Instruction


## Benchmarking Instruction

    go test -benchtime=10x -bench=BenchmarkSpade -benchmem -run=^$



```mermaid
sequenceDiagram
    autonumber
    actor u as User(s)|id
    participant c as Curator/Server
    actor a as Analyst(s)
    %% start the protocol 
    note over c: generate q, g <br> n=maximum number of users, <br> m=maximum number of elements

    c ->> c: msk, mpk = setup(n, m)
    c -->> u: (q, g, mpk)
    c -->> a: (q, g, mpk)

    u ->> u: generate random #alpha; <br> regKey = g^#alpha; <br> ct = Enc(mpk, #alpha;, m)
    u -->> c: (id, regKey, ct)

    c ->> c: stores regKeys[id] = regKey <br> stores cts[id] = ct

    %% analyst query
    a -->> c: (id, v)
    c ->> c: dkv = keyDer(id, v, msk, regKeys[id])
    c -->> a: (dkv, c[id])
    %% end of protocol
```
