using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class BattleMaster : MonoBehaviour
{
    [SerializeField] Bochamon _myBochamon;
    [SerializeField] Bochamon _enemyBochamon;

    [SerializeField] BattleUI _UI;

    public void Setup(Bochamon my, Bochamon enemy){
        _myBochamon = my;
        _enemyBochamon = enemy;

        _UI.Setup(_myBochamon, _enemyBochamon);
    }

    public void ExecuteMessage(){
        // Load battle - player1/player2, bochamon p1, 
        //  _uiController.InitialSetup
        
        // Change Bochamon -Player & bochamon
       //  _uiController.ChangeSelected

        // ExecuteMove - player - bochamon - move
        
        // out of battle
    }   

    void Update(){
        // Listen for server updates;
    }

    void SendMove(){
        Debug.Log("Send Move to server");
    }
}
