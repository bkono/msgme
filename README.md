# msgme

Sample app showing [fuse](https://www.fusetools.com) and [go-micro](https://github.com/micro) playing together in harmony, with a little [gomobile](https://github.com/golang/mobile) magic.

## What's the point?

One of my main side-projects is structured as an on-premise (in-home) set of core services, with many devices (tablets and others) relaying data and providing user interaction points. Working with Fuse has proven to be an absolutely awesome experience on the UI/UX front[^If you haven't looked into Fuse for your cross-platform mobile framework, you're missing out. Seriously, go look into it. This README can wait.]. Server-side, I love me some go-micro, and use it whenever I can. By bringing gomobile into the mix, I'm able to leverage the excellent client and streaming abstractions provided by go-micro, as well as the mdns registry for no-config discovery of services on a home network. This approach is night and day when compared to the very painful, manual approaches I've previously done just to get a single Android app discovering and interacting with simple gRPC services. 

## And why this repo?

I like to have reference repos. This is one of those. Specifically, it is put together to show a baseline config for all the parts working smoothly together.

## Getting things running

With a working environment, you can have this running in just a handful of build and deploy steps. Thanks Makefile!

### Pre-reqs

Make sure you have a working go, gomobile, protobuf, and fuse development environment.

### Build and Run

Build the protos. This step is optional, and will automatically be done by other steps. But its nice to triple check your proto environment is good to go.

```
make protos
```

In one terminal, run

```
make server
```

To confirm things are working correctly, open another terminal window and run the demo

```
make demo
```

Output should be something along the lines of:

![](https://dsh.re/b6a9f)

To explore the Fuse app, you'll need a connected android device. With the server still up and running, in another terminal, build the android sdk

```
make android
```

After a good bit of output, you can run the fuse preview

```
cd mobileapp
fuse preview
```

Once the fuse dummy app is loaded, you can either launch it from the app menu Preview -> Preview on Android, or via cli `fuse preview --target=android`. Either way, that step will take a moment, but eventually you'll have the app up and running on your device. Clicking listen will initiate a gRPC call via the gomobile sdk, which returns a stream of messages back to the device. These will be displayed (in a less than pretty way) as text bubbles in the UI.

### Acknowledgements

- Inspired by the [gomobile talk](https://youtu.be/LpWEDFT3iLU) from @sheriffjackson and the matching [sample code](https://github.com/gokitter).

