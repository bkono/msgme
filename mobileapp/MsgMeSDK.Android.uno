using Fuse;
using Fuse.Scripting;
using Uno.UX;
using Uno.Compiler.ExportTargetInterop;

[UXGlobalModule]
[Require("Gradle.Repository", "flatDir { dirs 'src/main/libs' }")]
[Require("Gradle.Dependency", "compile(name:'mobilesdk',ext:'aar')")]
public extern(Android) class MsgMeSDK : NativeEventEmitterModule {
  static readonly MsgMeSDK _instance;
  public static Java.Object _client;

  public MsgMeSDK() : base(true, "messageReceived") {
    if (_instance != null) return;

    _instance = this;
    Resource.SetGlobalKey(_instance, "MsgMeSDK");
    AddMember(new NativeFunction("Test", (NativeCallback)Test));
    AddMember(new NativeFunction("Send", (NativeCallback)Send));
    AddMember(new NativeFunction("Listen", (NativeCallback)Listen));

    CreateClient();
    if (_client == null) {
      debug_log "client is null, hmmmm";
    }

    SendImpl("startup", "init message");
  }

  static void OnMessageReceived(string message) {
    _instance.Emit("messageReceived", message);
  }

  static object Test(Context c, object[] args) {
    return "hi";
  }

  static object Send(Context c, object[] args) {
    if (args.Length != 2) {
        debug_log "only 2 args";
        return null;
    }

    debug_log args;
    SendImpl(args[0] as string, args[1] as string);

    return null;
  }

  static object Listen(Context c, object[] args) {
    ListenImpl();
    return null;
  }

  [Foreign(Language.Java)]
  public static void CreateClient() {
    @{
      msgmesdk.Client client = msgmesdk.Msgmesdk.newClient();
      @{_client:Set(client)};
    @}
  }

  [Foreign(Language.Java)]
  public static void SendImpl(string from, string message) {
    @{
      try {
        msgmesdk.Client client = (msgmesdk.Client) @{_client:Get()};
        client.send(from, message);
      } catch(Exception e) {
        e.printStackTrace(); 
      }
    @}
  }

  [Foreign(Language.Java)]
  public static void ListenImpl() {
    @{
      try {
        msgmesdk.Client client = (msgmesdk.Client) @{_client:Get()};
        msgmesdk.MessageCallback cb = new msgmesdk.MessageCallback() {
            @Override
            public void onMessage(final String from, final String message, final long timeEpoch) {
               @{OnMessageReceived(string):Call(message)};
            }
        };

        client.listen(cb);
      } catch(Exception e) {
        e.printStackTrace();
      }
    @}
  }
}