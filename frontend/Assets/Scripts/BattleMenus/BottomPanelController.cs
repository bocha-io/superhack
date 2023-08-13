using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;
using TMPro;
using DG.Tweening;

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
    Player _myself;

    public void InitialSetup(Player me, Bochamon _enemy){
        // Setup Actions?
        // Setup Moves
        // Setup bochamons
        SetupPlayer(me);
        StartCoroutine(SetupMyBochamon(me.bochamons[0]));
        StartCoroutine(SetupEnemyBochamon(_enemy));
    }

    public void ExitBattle(bool victory){
    }

    public void ForceBochamonChange(){
        _myBochamonSprite.gameObject.transform.DOLocalMoveY(-4f, 0.2f);
        _bochamonMenu.canGoBack = false;
        _bochamonMenu.Setup(_myself.bochamons);
        ChangeState(PanelState.BochamonMenu);
    }

    public void DefeatEnemyBochamon(){
        _enemyBochamonSprite.gameObject.transform.DOLocalMoveY(-4f, 0.2f);
    }

    public void SetupPlayer(Player player){    
        _myself = player;
        _bochamonMenu.Setup(player.bochamons);
    }

    public IEnumerator SetupMyBochamon(Bochamon bochi){
        // Setup UI
        if (_myBochamon.bochamon != null && bochi.uuid == _myBochamon.bochamon.uuid)
            yield break;
        ChangeState(PanelState.ShowingText);
        StartCoroutine(_battleInfo.ShowText("Sending out " + bochi.bochaName));
        yield return new WaitForSeconds(1.5f);
        _myBochamon.Setup(bochi);
        _fightMenu.Setup(bochi.moves);
        
        _myBochamonSprite.sprite = bochi.sprite;
        _myBochamonSprite.transform.position = new Vector3(-5.42f, 0.701f, _myBochamonSprite.transform.position.z);
        _myBochamonSprite.transform.DOLocalMoveX(_myBochamonSprite.transform.position.x+3, 0.5f);
        yield return new WaitForSeconds(1.5f);
        // Setup Moves 
    }

    public IEnumerator MyBochamonWasDefeated(){
        yield break;
    }

    public IEnumerator EnemyBochamonWasDefeated(){
        yield break;
    }

    public IEnumerator SetupEnemyBochamon(Bochamon bochi){
        // Setup UI
        if (_enemyBochamon.bochamon != null && bochi.uuid == _enemyBochamon.bochamon.uuid)
            yield break;
        ChangeState(PanelState.ShowingText);
        StartCoroutine(_battleInfo.ShowText("Enemy is Sending out " + bochi.bochaName));
        yield return new WaitForSeconds(1.5f);
        _enemyBochamon.Setup(bochi);
        _enemyBochamonSprite.sprite = bochi.sprite;
        _enemyBochamonSprite.transform.position = new Vector3(5.33f, 1.14f, _enemyBochamonSprite.transform.position.z);
        _enemyBochamonSprite.transform.DOLocalMoveX(_enemyBochamonSprite.transform.position.x-3, 0.5f);
        yield return new WaitForSeconds(1.5f);
    }

    public IEnumerator ApplyDamageOnSelf(int damage){
        ChangeState(PanelState.ShowingText);
        StartCoroutine(_battleInfo.ShowText("Enemy attacked us for " + damage.ToString() + " damage "));
        _myBochamonSprite.DOFade(0,0.1f).SetLoops(8, LoopType.Yoyo);
        yield return new WaitForSeconds(1.5f);
        _myBochamon.ApplyDamage(damage);
        yield return new WaitForSeconds(1.5f);
    }

    public IEnumerator ApplyDamageOnEnemy(int damage){
        ChangeState(PanelState.ShowingText);
        StartCoroutine(_battleInfo.ShowText("We caused " + damage.ToString() + " damage to enemy bochamon!"));
        _enemyBochamonSprite.DOFade(0,0.1f).SetLoops(8, LoopType.Yoyo);
        yield return new WaitForSeconds(1.5f);
        _enemyBochamon.ApplyDamage(damage);
        yield return new WaitForSeconds(1.5f);
    }

    public void ShowWaitingForOponnent(){
        ChangeState(PanelState.ShowingText);
        _battleInfo.ShowPermanentText("Waiting for opponent...");
    }

    public void ShowBattleResult(bool victory){
        ChangeState(PanelState.ShowingText);
        string text = "You lost the battle";
        if (victory){
            text = "You won the battle";
        }
        _battleInfo.ShowPermanentText(text);
    }

    public void OpenBochamon(){
        _bochamonMenu.Setup(_myself.bochamons);
        _bochamonMenu.canGoBack = true;
        ChangeState(PanelState.BochamonMenu);
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
                _bochamonMenu.Setup(_myself.bochamons);
                break;
            }
            default:
            {
                break;
            }
        }
    }


}
