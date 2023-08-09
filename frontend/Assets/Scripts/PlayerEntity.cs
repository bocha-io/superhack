using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using DG.Tweening;

public class PlayerEntity : MonoBehaviour
{
    (int x, int y) pos;
    string uuid;
    bool moving = false;
    [SerializeField] Animator _animator;
    public void Spawn(int x, int y, string id){
        pos = (x,y);
        uuid = id;
    }

    public void MoveTo(int x, int y, System.Action<int, int> action){
        if (moving) return;
        moving = true;
        int direction = GetMovementDirection(x, y);
        _animator.SetInteger("direction", direction);

        transform.DOMove(new Vector3(x, y, 1), 0.2f, false).OnStart(() => {
            action(x,y);
        }).OnComplete( () => {
            moving = false;
            _animator.SetInteger("direction", 0);
        });
        pos = (x,y);
    
        
        // if (Mathf.Abs(pos.x - x) + Mathf.Abs(pos.y - y) == 1){
        //     moving = true;
            
        // }//else transform.position = new Vector3(x,y,1);
        
    }

    int GetMovementDirection(int x, int y){
        if (x > pos.x) return 1;
        if (x < pos.x) return -1;
        if (y > pos.y) return 2;
        if (y < pos.y) return -2;
        return 0;
    }


}
