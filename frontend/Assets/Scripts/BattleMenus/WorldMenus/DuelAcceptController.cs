using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class DuelAcceptController : OptionsController
{    
    // public DuelAcceptOption yes;
    // public DuelAcceptOption no;

    public void Setup(string playerId, WorldCanvasController _parent){
        //
        //
        options[0].GetComponent<DuelAcceptAction>().player = playerId;
        options[1].GetComponent<DuelAcceptAction>().player = playerId;

        
    }



}
