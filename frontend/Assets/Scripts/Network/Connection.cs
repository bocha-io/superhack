using System;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;
using Newtonsoft.Json;
using NativeWebSocket;

using UnityEngine.SceneManagement;
using System.Runtime.InteropServices;

public class Connection : MonoBehaviour
{
  public static Connection Instance;

  NativeWebSocket.WebSocket websocket;
  public Queue<(string, string)> messages = new Queue<(string, string)>();

  #if UNITY_WEBGL && !UNITY_EDITOR
    [DllImport("__Internal")]
    private static extern string GetLocalStorage(string str);
  #endif

  public string url;
  
  public virtual void Awake()
  {
      // create the instance
      if (Instance == null)
      {
          DontDestroyOnLoad(this.gameObject);
      }
      else
      {
          Destroy(gameObject);
      }
  }

  async void Start()
  {
    string url2 = "";
    #if UNITY_WEBGL && !UNITY_EDITOR
      url2 = GetLocalStorage("url");
      Debug.Log("URL DEFINED IS " + url);
      url = url2=="" || url2 == null?url:url2;
    #endif

    websocket = new NativeWebSocket.WebSocket(url);

    websocket.OnOpen += () =>
    {
    };

    websocket.OnError += (e) =>
    {
      Debug.Log("Error! " + e);
    };

    websocket.OnClose += (e) =>
    {
    };

    websocket.OnMessage += (bytes) =>
    {
      // var message = System.Text.Encoding.UTF8.GetString(bytes);
      // // BaseMsg msg = JsonConvert.DeserializeObject<BaseMsg>(message);
      // messages.Enqueue((msg.msgtype, message));
    };

    // Keep sending messages at every 0.3s
    // InvokeRepeating("SendWebSocketMessage", 0.0f, 0.3f);

    // waiting for messages
    await websocket.Connect();
  }

  void Update()
  {
    #if !UNITY_WEBGL || UNITY_EDITOR
      websocket.DispatchMessageQueue();
    #endif
  }

  async public void SendWebSocketMessage(string message)
  {
    if (websocket.State == WebSocketState.Open)
    {
      await websocket.SendText(message);
    }
  }

  private void OnApplicationQuit()
  {
    Close();
  }

  public async void Close(){
      if (websocket != null && websocket.State == WebSocketState.Open)
        await websocket.Close();
  }

}