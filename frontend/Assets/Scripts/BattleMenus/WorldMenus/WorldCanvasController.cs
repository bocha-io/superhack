using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class WorldCanvasController : MonoBehaviour
{
    [SerializeField] DuelAcceptController _duel;
    [SerializeField] PlayerController _player;

    [SerializeField]  TradePokemonController _pokemonSelects;
    [SerializeField]  BochamonPanelController _trades;
    [SerializeField] GameObject _tradeCanvas;
    [SerializeField] Canvas _exchangeCanvas;

    public void ChallengedToDuel(string player){
        _player.inWorld = false;
        _duel.Setup(player, this);
        _duel.gameObject.SetActive(true);
    }

    public void CloseDuel(){
        _player.inWorld = true;
        _duel.gameObject.SetActive(false);
    }

    public void SelectChain(string chain){

    }

    public void OpenBochamonTrade(){
        _trades.enable = true;
        _trades.Setup(_player.bochamons);
        _trades.gameObject.SetActive(true);
        _tradeCanvas.SetActive(true);
        _pokemonSelects.gameObject.SetActive(false);
    }

    public void OpenExchange(){
        _player.inWorld = false;
        _exchangeCanvas.gameObject.SetActive(true);
    }

    public void Close(){
        _player.inWorld = true;

        _trades.gameObject.SetActive(false);
        _tradeCanvas.SetActive(false);
        _exchangeCanvas.gameObject.SetActive(false);
    }
}
