# meteor-fuzzy
A tool that reads and monitor data from Meteor's MongoDB database and turns them
into searchable data using fuzzy searching.

## Motivations

- Go memory usage is very appealing in most situations. Especially for my current
  usecase, other solutions like Elastic Search is just a bit overkill.
- Go is a compiled language, and generates a static binary that can run on any
  platform of the same architecture/system. One needs to compile the tool once
  and just copy it everywhere it is needed without worrying if it will work or
  not.
- The library `http://github.com/renstrom/fuzzysearch/fuzzy` is exactly what I
  needed for my current usecase.

## Installation
This tool uses ZeroMQ (ZMQ) version 4, which is not present on Ubuntu (don't know
what the reason behind thins). Thus, if you are on Ubuntu, even the latest one at
the time of writing this README is still version 2, and you will need to manually
install the libzmq version 4 manually.

### ZMQ Library for your System
This section is for installing the library itself, the core of ZMQ to be widely
available for your system. It is recommended to use container such as Docker so
that you don't pollute your system; I mean, if you think it could conflict with
the system package.

```sh
cd tmp/
wget 'http://download.zeromq.org/zeromq-4.1.4.tar.gz'
tar xf zeromq-4.1.4.tar.gz
cd zeromq-*
./configure --without-documentation --without-libsodium
make -j8
sudo make install
sudo ldconfig
```

### ZMQ Support for your Meteor app
Obviously, you need to be able to communicate with this tool in your Meteor app
as well. This is done via the wonderful `meteorhacks:npm` Meteor package.

```sh
cd meteor-app-directory/
meteor add meteorhacks:npm
```

You need to restart your Meteor app after the installation. Then, look into the
root of your Meteor app, there will be a new file called `packages.json`, open
and update it to contain this:

```json
{
  "zmq": "2.14.0"
}
```

At the time or writing this README, that is the most recent version of `zmq`
package for NPM. And yes, it is a NPM package; `meteorhacks:npm` is meant to
allow you to use NPM packages from within Meteor.

### Compilation
Since this is a Golang project, you need to compile it. One should be able to create the binary like:

```sh
mkdir $HOME/go
export GOPATH=$HOME/go
go get install github.com/eliezedeck/meteor-fuzzy
```

Now, in the directory `$HOME/go/bin` you will find a static binary file named `meteor-fuzzy`. That is your main binary, you can run in anywhere as long as the target platform is the same as yours.

## Usage
First, get documented on how to use ZeroMQ/ZMQ. Send your query request to the meteor-fuzzy as a JSON string with the following format:

```json
{
  "query": "your query here",
  "limit": 5
}
```

You will then receive a maximum amount of 5 (or any number you set). The format is:

```json
{
  "result": [
    "id1",
    "id2",
    "..."
  ]
}
```

Each of these are the ID of the items matched on the actual MongoDB database.

## Implementation
`meteor-fuzzy` makes use of the library http://github.com/renstrom/fuzzysearch/fuzzy
for the actual fuzzy searching algorithm. I had it forked to implement some of my requirements.

With that said, this project's is just to connect the dots. That is, inter-connect
Meteor's MongoDB and the actual algorithm that does the actual searching.

The Interconnection part is done thru the help of ZMQ library.

## LICENSE
MIT
