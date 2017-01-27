# Practical Golang: Bulding a simple distributed one-value database with Hashicorp Serf

## Introduction
With the advent of *distributed applications*, we see new storage solutions constantly.
They include, but are not limitted to, [Cassandra][1], [Redis][2], [CockroachDB][3], [Consul][4] or [RethinkDB][5].
Most of you probably use one, or more, of them. 

They seem to be really complex systems, because they actually are, which can't be denied. 
But it's pretty easy to write a simple, one value, database featuring *high availability*.
You probably wouldn't use anything near this in production, but it should be a fruitful learning experience for you nevertheless. 
If you're interested, read on!

## Dependencies
You'll need to
```
go get github.com/hashicorp/serf/serf
```
as a key dependency.

We'll also use those for convenience's sake:
```
"github.com/gorilla/mux"
"github.com/pkg/errors"
"golang.org/x/sync/errgroup"
```

## Small overview
What will we build? We'll build a one-value *clustered* database. Which means, numerous instances of our application will be able to work together.
You'll be able to set or get the value using a REST interface. The value will then shortly be spread across the cluster using the Gossip protocol. 
Which means, every node tells a part of the clsuter about the current state of the variable in set intevals. But because later each of those also tells a part of the cluster about the state, the whole cluster ends up having been informed. 

It'll use Serf for easy cluster membership, which uses SWIM under the hood. SWIM is a more advanced Gossipl-like algorithm, which you can read on about [here][6].

Now let's get to the implementation...

## Getting started

First we'll of course have to put in all our imports:
```go
import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/hashicorp/serf/serf"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)
```

Following this it's time to write a simple thread-safe one-value store.
An important thing is, the database will also hold the *generation* of the variable.
This way, when one instance gets notified about a new value, it can check if the incoming notification actually has a higher generation count.
So the our structure will hold exactly this: the number, generation and a mutex.

```go
type oneAndOnlyNumber struct {
	num        int
	generation int
	numMutex   sync.RWMutex
}

func InitTheNumber(val int) *oneAndOnlyNumber {
	return &oneAndOnlyNumber{
		num: val,
	}
}
```

We'll also need a way to set and get the value. 
Setting the value will also advance the generation count, so when we notify the rest of this cluster, we will overwrite their values and generation counts.

```go
func (n *oneAndOnlyNumber) setValue(newVal int) {
	n.numMutex.Lock()
	defer n.numMutex.Unlock()
	n.num = newVal
	n.generation = n.generation + 1
}

func (n *oneAndOnlyNumber) getValue() (int, int) {
	n.numMutex.RLock()
	defer n.numMutex.RUnlock()
	return n.num, n.generation
}
```

Finally, we will need a way to notify the database of changes that happened elsewhere, if they have a higher generation count.
For that we'll have a small notify method, which will return true, if anything has been changed:

```go
func (n *oneAndOnlyNumber) notifyValue(curVal int, curGeneration int) bool {
	if curGeneration > n.generation {
		n.numMutex.Lock()
		defer n.numMutex.Unlock()
		n.generation = curGeneration
		n.num = curVal
		return true
	}
	return false
}
```

We'll also create a const describing how many nodes we will notify about the new value every time.

```go
const MembersToNotify = 2
```

Now let's get to the actual functioning of the application. First we'll have to start an instance of serf, using two variables. The address of our instance in the network and the -optional- address of the cluster to joing.

```go
func main() {
	cluster, err := setupCluster(
		os.Getenv("ADVERTISE_ADDR"),
		os.Getenv("CLUSTER_ADDR"))
	if err != nil {
		log.Fatal(err)
	}
	defer cluster.Leave()
```

How does the setupCluster function work, you may ask? Here it is:

```go
func setupCluster(advertiseAddr string, clusterAddr string) (*serf.Serf, error) {
	conf := serf.DefaultConfig()
	conf.Init()
	conf.MemberlistConfig.AdvertiseAddr = advertiseAddr

	cluster, err := serf.Create(conf)
	if err != nil {
		return nil, errors.Wrap(err, "Couldn't create cluster")
	}

	_, err = cluster.Join([]string{clusterAddr}, true)
	if err != nil {
		log.Printf("Couldn't join cluster, starting own: %v\n", err)
	}

	return cluster, nil
}
```


As we can see, we are creating the cluster, only changing the advertise address.

If the creation fails, we return the error of course.
If the joining fails though, it means that we either didn't get a cluster address, or the cluster doesn't exist (omitting network failures), which means we can safely ignore that and just log it.




[1]:https://cassandra.apache.org/
[2]:https://redis.io/
[3]:https://www.cockroachlabs.com/
[4]:https://www.consul.io/
[5]:https://www.rethinkdb.com/
[6]:https://www.cs.cornell.edu/~asdas/research/dsn02-swim.pdf