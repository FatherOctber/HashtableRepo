# HashtableRepo
:beers: Hashtable repository in a multithreaded environment (Golang)

## Description:
Implementation of synchronized operation of external clients with the data store (server) presented in the hashtable form.

## Conditions:
- Multiple clients can simultaneously read from a remote storage;  
- Writing in the store at the same time produces only one client;  
- The implementation of a mechanism to allow the transaction "rollback" of changes in the event of an error during a write operation;  
- Network service: [XML-RPC](https://github.com/divan/gorilla-xmlrpc/).  

##### For the hashtable the following methods are realized:
```go
bool clear();
bool containsKey(int key);
bool containsValue(string value);
string get(int key);
bool isEmpty();
string put(int key, string value);
string remove(int key);
int size();
int hashcode(int key);
string toString();
```

##### Also, the hashtable has the following properties:
```java
Entry[] table;
float loadFactor;
int capacity;
int size;
int threshold;
```
