using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using Newtonsoft.Json;

public class DuelAcceptAction : MenuOptionAction
{
    public string player;
    public bool accept;
    [SerializeField] WorldCanvasController _parent;

    public override void Execute()
    {
        // Connection.Instance.
        DuelResponseMessage response = new(){
            msgtype = "duelresponse",
            response = accept,
            playera = player
        };
        Connection.Instance.SendWebSocketMessage(JsonConvert.SerializeObject(response));
        _parent.CloseDuel();
    }
}
