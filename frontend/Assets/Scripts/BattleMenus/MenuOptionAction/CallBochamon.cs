using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class CallBochamon : MenuOptionAction
{
    [SerializeField] BattleMaster _battleMaster;
    [SerializeField] BochamonOption _bochaOption;
    public int pos;
    public override void Execute(){
        if (_bochaOption._bochamon.currentHp == 0)
             return;
        _battleMaster.SendBochamon(pos);
        Debug.Log("Send Bochamon " + _bochaOption._bochamon.bochaName);
    }
}
