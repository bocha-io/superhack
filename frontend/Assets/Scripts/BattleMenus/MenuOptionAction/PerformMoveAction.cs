using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class PerformMoveAction : MenuOptionAction
{
    [SerializeField] BattleMaster _battleMaster;
    public int number;
    public override void Execute(){
        // _battleMaster.
        _battleMaster.SendMove(number);        
        Debug.Log("Send Action " + number);
    }
}
