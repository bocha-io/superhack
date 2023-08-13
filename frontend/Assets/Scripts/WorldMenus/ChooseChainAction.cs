using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class ChooseChainAction : MenuOptionAction
{
    [SerializeField] WorldCanvasController _panels;
    public string chain;

    public override void Execute()
    {
        _panels.SelectChain(chain);
    }
}
