using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class RunOption : MenuOptionAction
{
    [SerializeField] BattleMaster _battleMaster;
    public override void Execute(){
        // _battleMaster.
        Debug.Log("Exit Battle");
    }
}
