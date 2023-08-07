using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class OpenFight : MenuOptionAction
{
    [SerializeField] BottomPanelController _battleUI;

    public override void Execute(){
        _battleUI.ChangeState(PanelState.FightMenu);
    }


}
