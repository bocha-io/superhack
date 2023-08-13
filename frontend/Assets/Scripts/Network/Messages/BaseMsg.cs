using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class BaseMsg 
{
    public string msgtype;
}

public class PlayerPos
{
    public int X;
    public int Y;
    public string playerid;
}


public class MapStatus : BaseMsg
{
    public PlayerPos[] playerspos;
    /*msgtype = mapstatus*/
}

public class MoveMessage : BaseMsg
{
    /*msgtype = move*/
    public int x;
    public int y;
}

public class MoveResponse : BaseMsg
{
    /*msgtype  = moveresponse */
    public bool value;
    public string error;
}

public class Match 
{
    public string id;
    public string playerone;
    public string playertwo;
}

public class CreateMatchMessage : BaseMsg
{
    /*msgtype creatematch*/
    public string playerA;
    public string playerB;
}

public class CreateMatchMessageResponse : BaseMsg
{    
    /*creatematchmessageresponse*/
    public Match value;
    public string error;
}

public class BochamonMsg
{
    public string id;
    public int hp;
    public int montype;
}

public class BochamonsMsg
{
    public BochamonMsg first;
    public BochamonMsg second;
    public BochamonMsg third;
}


/*
{
   "msgtype":"battlestatus",
   "playerone":"0x6640fa6e1fac1a10d5235196e7e467674cf36c96",
   "playertwo":"0xcc9767009271154bfa386b3a317573a9d99d1ca9",
   "matchid":"0x0000000000000000000000006640fa6e1fac1a10d5235196e7e467674cf36c96",
   "actions":{
      "playeroneswapped":false,
      "playertwoswapped":false,
      "damagedunits":null,
      "playeroneattack":0,
      "playertwoattack":0,
      "winner":""
   },
   "playeronemons":{
      "first":{
         "id":"0x3100000000000000000000006640fa6e1fac1a10d5235196e7e467674cf36c96",
         "hp":450,
         "montype":1
      },
      "second":{
         "id":"0x3200000000000000000000006640fa6e1fac1a10d5235196e7e467674cf36c96",
         "hp":500,
         "montype":3
      },
      "third":{
         "id":"0x3300000000000000000000006640fa6e1fac1a10d5235196e7e467674cf36c96",
         "hp":400,
         "montype":6
      }
   },
   "playertwomons":{
      "first":{
         "id":"0x310000000000000000000000cc9767009271154bfa386b3a317573a9d99d1ca9",
         "hp":450,
         "montype":1
      },
      "second":{
         "id":"0x320000000000000000000000cc9767009271154bfa386b3a317573a9d99d1ca9",
         "hp":500,
         "montype":3
      },
      "third":{
         "id":"0x330000000000000000000000cc9767009271154bfa386b3a317573a9d99d1ca9",
         "hp":400,
         "montype":6
      }
   },
   "playeronecurrentmon":"0x3100000000000000000000006640fa6e1fac1a10d5235196e7e467674cf36c96",
   "playertwocurrentmon":""
}
*/

public class Actions 
{
    public bool playeroneswapped;
    public bool playertwoswapped;
    public string[] damagedunits;
    public int playeroneattack;
    public int playertwoattack;
    public string winner;
}
public class BattleStatus : BaseMsg 
{
    /*msgtype battlestatus*/
    public string playerone;
    public string playertwo;
    public string matchid;

    public Actions actions;

    public BochamonsMsg playeronemons;
    public BochamonsMsg playertwomons;

    public string playeronecurrentmon;
    public string playertwocurrentmon;
}

public class BattleAction : BaseMsg
{
    public string matchid;
    public int playeroneaction;
    public int playeroneactionpos;

    public int playertwoaction;
    public int playertwoactionpos;
}

public class ConnectMsg : BaseMsg
{
    public string user;
    public string password;
}

public class ConnectResponse : BaseMsg
{
    public string value;
    public string error;
}


public class DuelRequestMessage : BaseMsg
{
    public string enemy;
}

public class Duel
{
    public string playera;
    public string playerb;
}

public class DuelRequestMessageResponse : BaseMsg
{
    public Duel value;
    public string error;
}

public class DuelResponseMessage : BaseMsg
{
    public bool response;
    public string playera;
}

public class DuelResponseMessageResponse : BaseMsg
{
    public Duel value;
    public string error;
}

public class SendAction : BaseMsg
{
    public int action;
    public int pos;
    public string matchid;
}


public class SwapMessage : BaseMsg 
{
    public int montype;
    public int pos;
}

public class SwapMessageResponse : BaseMsg
{
    public bool value;
    public string error;
}

public class InventoryResponse : BaseMsg
{
    public int[] value;
    public string error;
}

public class BridgeMessage : BaseMsg
{
    public int amount;
}

public class BridgeMessageResponse : BaseMsg
{
    public string value;
    public string error;
}