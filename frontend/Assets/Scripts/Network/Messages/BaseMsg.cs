using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class BaseMsg 
{
    public string msgtype;
}

public class PlayerPos
{
    public int x;
    public int y;
    public string id;
}


public class MapStatus : BaseMsg
{
    public PlayerPos[] playerpos;
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

public class BattleStatus : BaseMsg 
{
    /*msgtype battlestatus*/
    public string playerone;
    public string playertwo;
    public string matchid;

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
