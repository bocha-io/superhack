using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class BochamonTradeController : BochamonPanelController
{
    [SerializeField] WorldCanvasController _worldUI;
    public override void GoBack(){
        _worldUI.Close();
    }
}
