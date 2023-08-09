using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;
using TMPro;

public enum PanelState {
    ShowingText,
    PickingAction,
    FightMenu,
    BochamonMenu
}

public class BottomPanelController : MonoBehaviour
{
    [SerializeField] OptionsController _actions;
    [SerializeField] MovesPanelController _fightMenu;    
    [SerializeField] BochamonPanelController _bochamonMenu;
    [SerializeField] BattleInfo _battleInfo;

    [SerializeField] BochamonUI _myBochamon;
    [SerializeField] BochamonUI _enemyBochamon;

    [SerializeField] SpriteRenderer _myBochamonSprite;
    [SerializeField] SpriteRenderer _enemyBochamonSprite;

    public PanelState currentState;

    public void InitialSetup(Player me, Bochamon _enemy){
        // Setup Actions?
        // Setup Moves
        // Setup bochamons
        SetupPlayer(me);
        SetupMyBochamon(me.bochamons[0]);
        SetupEnemyBochamon(_enemy);
    }

    public void SetupPlayer(Player player){    
        _bochamonMenu.Setup(player.bochamons);
    }

    public void SetupMyBochamon(Bochamon bochi){
        // Setup UI
        _myBochamon.Setup(bochi);
        _fightMenu.Setup(bochi.moves);
        _myBochamonSprite.sprite = bochi.sprite;
        // Setup Moves 
    }

    public void SetupEnemyBochamon(Bochamon bochi){
        // Setup UI
        _enemyBochamon.Setup(bochi);
        _enemyBochamonSprite.sprite = bochi.sprite;
    }

    public void ApplyDamageOnSelf(int damage){
        _myBochamon.ApplyDamage(damage);
    }

    public void ApplyDamageOnEnemy(int damage){
        _enemyBochamon.ApplyDamage(damage);
    }

    public void ChangeState(PanelState state){
        _battleInfo.gameObject.SetActive(state == PanelState.ShowingText);
        _actions.gameObject.SetActive(state == PanelState.PickingAction);
        _fightMenu.gameObject.SetActive(state == PanelState.FightMenu);
        _bochamonMenu.gameObject.SetActive(state == PanelState.BochamonMenu);

        switch(currentState){
            case PanelState.ShowingText:
            {

                break;
            }
            case PanelState.PickingAction:
            {
                break;
            }
            case PanelState.FightMenu:
            {
                break;
            }
            case PanelState.BochamonMenu:
            {
                break;
            }
            default:
            {
                break;
            }
        }
    }


}
