using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using DG.Tweening;

public class PlayerEntity : MonoBehaviour
{
    (int x, int y) pos;
    public string uuid;
    bool moving = false;
    [SerializeField] public Animator _animator;
    public void Spawn(int x, int y, string id){
        pos = (x,y);
        uuid = id;
    }

    public void SetPosition(int x, int y){
        pos = (x,y);
        transform.position = new Vector3(x,y,1);
    }
    public int MoveTo(int x, int y, System.Action<int, int> action){
        if (moving) return 0;
        moving = true;
        int direction = GetMovementDirection(x, y);
        _animator.SetInteger("direction", direction);
        _animator.SetBool("moving", true);
        
        int amountMoved = 1; //Mathf.Abs(x-pos.x) + Mathf.Abs(y-pos.y);

        transform.DOMove(new Vector3(x, y, 1), 0.33f * amountMoved, false).OnStart(() => {
            action(x,y);
        }).OnComplete( () => {
            moving = false;
            _animator.SetInteger("direction", 0);
            _animator.SetBool("moving", false);
        });
        pos = (x,y);
    
        return direction;
        // if (Mathf.Abs(pos.x - x) + Mathf.Abs(pos.y - y) == 1){
        //     moving = true;
            
        // }//else transform.position = new Vector3(x,y,1);
        
    }

    public int GetMovementDirection(int x, int y){
        if (x > pos.x) return 1;
        if (x < pos.x) return -1;
        if (y > pos.y) return 2;
        if (y < pos.y) return -2;
        return 0;
    }


}
