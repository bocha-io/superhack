using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class MovesPanelController : OptionsController
{


    public override void Select(){
        for (int i=0; i< options.Count; i++){
            ((FightOption)options[i]).Select(i==selected);
        }
    }
    public void Setup(List<Moves> _moves){
        for(int i=0; i< options.Count; i++){
            ((FightOption)options[i]).Setup(_moves[i]);
        }
    }
}
