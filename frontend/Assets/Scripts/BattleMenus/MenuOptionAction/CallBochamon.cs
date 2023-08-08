using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class CallBochamon : MenuOptionAction
{
    [SerializeField] BattleMaster _battleMaster;
    public override void Execute(){
        // _battleMaster.
        Debug.Log("Send Bochamon");
    }
}
