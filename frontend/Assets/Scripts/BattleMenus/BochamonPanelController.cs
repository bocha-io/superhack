using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class BochamonPanelController : OptionsController
{

    public override void Select(){
        for (int i=0; i< options.Count; i++){
            ((BochamonOption)options[i]).Select(i==selected);
        }
    }
    public void Setup(List<Bochamon> bochamons){
        for(int i=0; i< options.Count; i++){
            ((BochamonOption)options[i]).Setup(bochamons[i]);
        }
    }

}
