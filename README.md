# meteor-fuzzy
A tool that reads and monitor data from Meteor's MongoDB database and turns them into searchable data using fuzzy searching.

## Installation
This tool uses ZeroMQ (ZMQ) version 4, which is not present on Ubuntu (don't know
what the reason behind thins). Thus, if you are on Ubuntu, even the latest one at
the time of writing this README is still version 2, and you will need to manually
install the libzmq version 4 manually.

```sh
cd tmp/
wget 'http://download.zeromq.org/zeromq-4.1.4.tar.gz'
tar xf zeromq-4.1.4.tar.gz
cd zeromq-*
./configure --without-documentation --without-libsodium
make -j8
sudo make install
```

## Implementation
`meteor-fuzzy` makes use of the library http://github.com/renstrom/fuzzysearch/fuzzy
for the actual fuzzy searching algorithm.

With that said, this project's is just to connect the dots. That is, inter-connect
Meteor's MongoDB and the actual algorithm that does the actual searching.
