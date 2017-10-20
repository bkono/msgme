using Fuse;
using Fuse.Scripting;
using Uno;
using Uno.UX;
using Uno.Threading;
using Uno.Compiler.ExportTargetInterop;

[UXGlobalModule]
public extern(!Android) class MsgMeSDK : NativeEventEmitterModule {
  static readonly MsgMeSDK _instance;
  public MsgMeSDK() : base(true, "messageReceived") {
    if (_instance != null) return;

    _instance = this;
    Resource.SetGlobalKey(_instance, "MsgMeSDK");
    AddMember(new NativeFunction("Test", (NativeCallback)Test));
    AddMember(new NativeFunction("Send", (NativeCallback)Send));
    AddMember(new NativeFunction("Listen", (NativeCallback)Listen));
  }

  static string Test(Context c, object[] args) {
    return "hi";
  }

  static object Listen(Context c, object[] args) {
    ListenImpl();
    return null;
  }

  static object Send(Context c, object[] args) {
    debug_log "would send message here";
    return null;
  }

  static void ListenImpl() {
    new Thread(EmitMessage).Start();
  }

  static void EmitMessage() {
    Thread.Sleep(1000);
    _instance.Emit("messageReceived", "some message emitted #1");
    Thread.Sleep(2000);
    _instance.Emit("messageReceived", "some message emitted #2");
    Thread.Sleep(1500);
    _instance.Emit("messageReceived", "some message emitted #3");
  }
}