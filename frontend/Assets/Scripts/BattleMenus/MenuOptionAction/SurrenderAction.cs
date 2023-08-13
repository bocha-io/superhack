using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class SurrenderAction : MenuOptionAction
{
    [SerializeField] BattleMaster _battleMaster;
    public override void Execute(){
        // _battleMaster.
        _battleMaster.SendSurrender();        
        Debug.Log("Send surrender ");
    }
}
