using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using Newtonsoft.Json;
using UnityEngine.SceneManagement;

public class BattleMaster : MonoBehaviour
{
    [SerializeField] Player _myself;
    [SerializeField] Player _enemy;

    [SerializeField] Bochamon _myBochamon;
    [SerializeField] Bochamon _enemyBochamon;

    [SerializeField] BottomPanelController _UI;

    [SerializeField] List<Bochamon> _bochamonsPrefabs;

    [SerializeField] TimerController _timer;

    string matchid = "";
    (string, string) message;

    bool imPlayerone = true;

    public void Start(){
        // Setup(_myself, _enemy);

        Debug.Log(PublicBattleStatus.valid);
        Debug.Log(PublicBattleStatus.status);
        if (PublicBattleStatus.valid){
            InitialSetup(PublicBattleStatus.status);
        }
    }

    public void InitialSetup(BattleStatus battle){

        
        Debug.Log("Seting Up battle" + JsonConvert.SerializeObject(battle));
        matchid = battle.matchid;
        
        _myself.bochamons.Clear();
        _enemy.bochamons.Clear();

        if(battle.playerone == Connection.Instance.wallet){
            imPlayerone = true;
            _myself.id = battle.playerone;
            _enemy.id = battle.playertwo;

            var playermons = battle.playeronemons;
            var player = _myself;
            for (int i=0; i<3; i++){
                var calling = i==0?playermons.first : (i==1)? playermons.second:playermons.third;
                Bochamon b1 = Instantiate(_bochamonsPrefabs[calling.montype], player.transform);
                b1.uuid = calling.id;
                b1.currentHp = calling.hp;
                player.bochamons.Add(b1);
            }
            
            playermons = battle.playertwomons;
            player = _enemy;
            for (int i=0; i<3; i++){
                var calling = i==0?playermons.first : (i==1)? playermons.second:playermons.third;
                Bochamon b1 = Instantiate(_bochamonsPrefabs[calling.montype], player.transform);
                b1.uuid = calling.id;
                b1.currentHp = calling.hp;
                player.bochamons.Add(b1);
            }
            Setup(_myself, _enemy);
        } else {
            Debug.Log("Im Player two");
            imPlayerone = false;
            _myself.id = battle.playertwo;
            _enemy.id = battle.playerone;

            var playermons = battle.playeronemons;
            var player = _enemy;
            for (int i=0; i<3; i++){
                var calling = i==0?playermons.first : (i==1)? playermons.second:playermons.third;
                Bochamon b1 = Instantiate(_bochamonsPrefabs[calling.montype], player.transform);
                b1.uuid = calling.id;
                b1.currentHp = calling.hp;
                player.bochamons.Add(b1);
            }
            
            playermons = battle.playertwomons;
            player = _myself;
            for (int i=0; i<3; i++){
                var calling = i==0?playermons.first : (i==1)? playermons.second:playermons.third;
                Bochamon b1 = Instantiate(_bochamonsPrefabs[calling.montype], player.transform);
                b1.uuid = calling.id;
                b1.currentHp = calling.hp;
                player.bochamons.Add(b1);
            }
            Setup(_myself, _enemy);
        }
        _timer.Reset(true);
    }

    public void Setup(Player me, Player enemy){
        _myself = me;
        _myBochamon = me.bochamons[0];
        _enemyBochamon = _enemy.bochamons[0];
        _enemy = enemy;

        _UI.InitialSetup(me, _enemyBochamon);
    }

    public void ExecuteMessage((string type, string content) message)
    {
        Debug.Log("Battle message");
        switch(message.type)
        {
            case "battlestatus":
            {
                BattleStatus battleStatus = JsonConvert.DeserializeObject<BattleStatus>(message.content);

                if (isInitialAction(battleStatus.actions))
                    InitialSetup(battleStatus);
                else 
                    StartCoroutine(ExecuteBattleActions(battleStatus));
                break;
            }
            default:
            {
                Debug.LogWarning("Message not recognized: " + message.type);
                break;
            }
        }
    }

    public bool isInitialAction(Actions action){
        if (!action.playeroneswapped && !action.playeroneswapped && action.damagedunits==null && action.winner==""){
            Debug.Log("IS INITIAL ACTION");
            return true;
        }
        Debug.Log("IS BATTLE ACTION");
        return false;
    }
    public IEnumerator ExitBattle(bool victory){
        // TODO: exit gracefully
        _UI.ShowBattleResult(victory);
        yield return new WaitForSeconds(4);
        SceneManager.LoadScene("WorldScene");
    }
 

    // Esperar entre accion y accion
    // Detectar cuando alguien muere y forzar el cambio
    // Detectar cuando el enemigo muere y automandar swap por self
    // Esperar al server

    // Marcar como muertos a los bochamones sin hp.

    IEnumerator ExecuteBattleActions(BattleStatus battleAction){
        Actions actions = battleAction.actions;

        if (actions.winner != null && actions.winner != ""){
            string temp = actions.winner.TrimStart('0');
            temp = temp.TrimStart('x');
            temp = temp.TrimStart('0');
            temp = "0x" + temp;
            StartCoroutine(ExitBattle(temp == Connection.Instance.wallet));
            yield break;
        }

        _timer.Reset(true);

        if(actions.playeroneswapped && imPlayerone){
            _myBochamon = _myself.GetBochamon(battleAction.playeronecurrentmon);
            yield return StartCoroutine(_UI.SetupMyBochamon(_myBochamon));
        }
        if(actions.playertwoswapped && imPlayerone){
            _enemyBochamon = _enemy.GetBochamon(battleAction.playertwocurrentmon);
            yield return StartCoroutine(_UI.SetupEnemyBochamon(_enemyBochamon));
        }

        if(actions.playeroneswapped && !imPlayerone){
            _enemyBochamon = _enemy.GetBochamon(battleAction.playeronecurrentmon);
            yield return StartCoroutine(_UI.SetupEnemyBochamon(_enemyBochamon));
        }
        if(actions.playertwoswapped && !imPlayerone){
            _myBochamon = _myself.GetBochamon(battleAction.playertwocurrentmon);
            yield return StartCoroutine(_UI.SetupMyBochamon(_myBochamon));
        }
        int myAttack, enemyAttack;
        BochamonsMsg myBochamons, enemyBochamons;
        if (imPlayerone){
            myAttack = actions.playeroneattack;
            enemyAttack = actions.playertwoattack;
            myBochamons = battleAction.playeronemons;
            enemyBochamons = battleAction.playertwomons;
        } else {
            myAttack = actions.playertwoattack;
            enemyAttack = actions.playeroneattack;
            myBochamons = battleAction.playertwomons;
            enemyBochamons = battleAction.playeronemons;
        } 


        if (actions.damagedunits != null){
            foreach(string unit in actions.damagedunits){
                if (unit == _myBochamon.uuid){
                    var bochamonData = GetBochamonData(unit, myBochamons);
                    var damage =  _myBochamon.currentHp - bochamonData.hp;
                    // -= _enemyBochamon.moves[enemyAttack].power;
                    _myBochamon.currentHp = bochamonData.hp;
                    yield return StartCoroutine(_UI.ApplyDamageOnSelf(damage));
                    if (bochamonData.hp == 0){
                        _UI.ForceBochamonChange();
                    }
                } else if (unit == _enemyBochamon.uuid){
                    var bochamonData = GetBochamonData(unit, enemyBochamons);
                    var damage =  _enemyBochamon.currentHp - bochamonData.hp;
                    // _enemyBochamon.currentHp -= _myBochamon.moves[myAttack].power;
                    _enemyBochamon.currentHp = bochamonData.hp;
                     yield return StartCoroutine(_UI.ApplyDamageOnEnemy(damage));
                    if (bochamonData.hp == 0){
                        _UI.DefeatEnemyBochamon();
                        SendBochamon(GetBochamonId(_myBochamon.uuid, myBochamons));
                        _UI.ShowWaitingForOponnent();
                    }
                } else {
                    Debug.LogError("Unrecognized bochamon");
                }
            }
        }
    }

    void PassTurn(){
        
    }

    public BochamonMsg GetBochamonData(string id, BochamonsMsg bochamons){
        if (bochamons.first.id == id)
            return bochamons.first;
        else if (bochamons.second.id == id)
            return bochamons.second;
        else if (bochamons.third.id == id)
            return bochamons.third;
        return null;
    }

    public int GetBochamonId(string id, BochamonsMsg bochamons){
        if (bochamons.first.id == id)
            return 0;
        else if (bochamons.second.id == id)
            return 1;
        else if (bochamons.third.id == id)
            return 2;
        return -1;
    }


    public void Update()
    {
        if(Connection.Instance.messages.TryDequeue(out message)){
            ExecuteMessage(message);
        }
    }

    public void SendSurrender(){
        SendAction action = new(){
            msgtype = "sendaction",
            action = 2,
            pos = 0,
            matchid = matchid
        };
        Connection.Instance.SendWebSocketMessage(JsonConvert.SerializeObject(action));
    }

    public void SendBochamon(int move){
        SendAction action = new(){
            msgtype = "sendaction",
            action = 1,
            pos = move,
            matchid = matchid
        };
        Connection.Instance.SendWebSocketMessage(JsonConvert.SerializeObject(action));
        _UI.ShowWaitingForOponnent();
    }
    
    public void SendMove(int move){
        SendAction action = new(){
            msgtype = "sendaction",
            action = 0,
            pos = move,
            matchid = matchid
        };
        Connection.Instance.SendWebSocketMessage(JsonConvert.SerializeObject(action));
        _UI.ShowWaitingForOponnent();
    }
}
