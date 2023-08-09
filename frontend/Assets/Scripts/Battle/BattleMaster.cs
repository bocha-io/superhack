using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using Newtonsoft.Json;

public class BattleMaster : MonoBehaviour
{
    [SerializeField] Player _myself;
    [SerializeField] Player _enemy;

    [SerializeField] Bochamon _myBochamon;
    [SerializeField] Bochamon _enemyBochamon;

    [SerializeField] BottomPanelController _UI;

    [SerializeField] List<Bochamon> _bochamonsPrefabs;

    public void Start(){
        Setup(_myself, _enemy);
    }

    public void InitialSetup(BattleStatus battle){
        _myself.id = battle.playerone;
        _enemy.id = battle.playertwo;
        _myself.bochamons.Clear();
        _enemy.bochamons.Clear();

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
        // Load battle - player1/player2, bochamon p1, 
        //  _uiController.InitialSetup
        
        // Change Bochamon -Player & bochamon
       //  _uiController.ChangeSelected

        // ExecuteMove - player - bochamon - move
        
        // out of battle
 
        switch(message.type)
        {
            case "battle":
            {
                BattleAction battleAction = JsonConvert.DeserializeObject<BattleAction>(message.content);
                ExecuteBattleActions(battleAction);
                break;
            }
            case "battlestatus":
            {
                BattleStatus battleStatus = JsonConvert.DeserializeObject<BattleStatus>(message.content);
                InitialSetup(battleStatus);
                break;
            }
            case "moveresponse":
            {
                // Verify movement
                break;
            }
            default:
            {
                Debug.LogWarning("Message not recognized: " + message.type);
                break;
            }
        }
    }
 
    void ExecuteBattleActions(BattleAction battleAction){
    }

    void ApplyDamage(){

    }

    void Update(){
        // Listen for server updates;
    }

    void SendMove(){
        Debug.Log("Send Move to server");
    }
}
