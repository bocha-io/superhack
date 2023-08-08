using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using DG.Tweening;

public class PlayerEntity : MonoBehaviour
{
    (int x, int y) pos;
    string uuid;
    bool moving = false;

    public void Spawn(int x, int y, string id){
        pos = (x,y);
        uuid = id;
    }

    public void MoveTo(int x, int y){
        if (moving) return;
        moving = true;
        transform.DOMove(new Vector3(x, y, 1), 0.2f, false).OnStart(() => {
    
        }).OnComplete( () => {
            moving = false;
        });
        pos = (x,y);
    
        
        // if (Mathf.Abs(pos.x - x) + Mathf.Abs(pos.y - y) == 1){
        //     moving = true;
            
        // }//else transform.position = new Vector3(x,y,1);
        
    }


}
