using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class ExchangeBochamonAction : MenuOptionAction
{
    public int montype;
    [SerializeField] TradePokemonController _trade;
    [SerializeField] WorldCanvasController _panel;
    public override void Execute()
    {
        _trade.PerformSwap(montype);
        _panel.Close();
    }
}
