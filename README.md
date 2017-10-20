# msgme

Sample app showing [fuse](https://www.fusetools.com) and [go-micro](https://github.com/micro) playing together in harmony, with a little [gomobile](https://github.com/golang/mobile) magic.

### What's the point?

One of my main side-projects is structured as an on-premise (in-home) set of core services, with many devices (tablets and others) relaying data and providing user interaction points. Working with Fuse has proven to be an absolutely awesome experience on the UI/UX front[^If you haven't looked into Fuse for your cross-platform mobile framework, you're missing out. Seriously, go look into it. This README can wait.]. Server-side, I love me some go-micro, and use it whenever I can. By bringing gomobile into the mix, I'm able to leverage the excellent client and streaming abstractions provided by go-micro, as well as the mdns registry for no-config discovery of services on a home network. This approach is night and day when compared to the very painful, manual approaches I've previously done just to get a single Android app discovering and interacting with simple gRPC services. 

### And why this repo?

I like to have reference repos. This is one of those. Specifically, it is put together to show a baseline config for all the parts working smoothly together.

