using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using Newtonsoft.Json;

public class SendExchangeAction : MenuOptionAction
{
    [SerializeField] WorldCanvasController _worldUI;

    public override void Execute()
    {
        BridgeMessage bridge = new(){
            msgtype="bridge",
            amount=1
        };
        Connection.Instance.SendWebSocketMessage(JsonConvert.SerializeObject(bridge));
        _worldUI.Close();
    }
}
