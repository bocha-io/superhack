using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class OpenBochamon : MenuOptionAction
{
    [SerializeField] BottomPanelController _battleUI;
    public override void Execute(){
        _battleUI.OpenBochamon();
        // _battleUI.ChangeState(PanelState.BochamonMenu);
    }
}
