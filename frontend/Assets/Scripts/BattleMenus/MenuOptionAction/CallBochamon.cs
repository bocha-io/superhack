using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class CallBochamon : MenuOptionAction
{
    [SerializeField] BattleMaster _battleMaster;
    [SerializeField] BochamonOption _bochaOption;
    public override void Execute(){
        // _battleMaster.
        Debug.Log("Send Bochamon " + _bochaOption._bochamon.bochaName);
    }
}
