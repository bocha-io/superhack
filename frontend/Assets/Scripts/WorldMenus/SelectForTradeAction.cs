using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class SelectForTradeAction : MenuOptionAction
{
    public int number;
    [SerializeField] TradePokemonController _panel;
    [SerializeField] BochamonPanelController _bochamonPanel;
    [SerializeField] PlayerController _player;

    public override void Execute()
    {
        _bochamonPanel.enable = false;
        _panel.gameObject.SetActive(true);
        _panel.Setup(_player.bochamons, number);
    }
}
