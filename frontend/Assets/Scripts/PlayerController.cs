using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class PlayerController : MonoBehaviour
{
    [SerializeField] PlayerEntity _player;
    [SerializeField] Collider2D _colliderUp;
    [SerializeField] Collider2D _colliderDown;
    [SerializeField] Collider2D _colliderLeft;
    [SerializeField] Collider2D _colliderRight;
    LayerMask mask;
    void Start(){
         mask = LayerMask.GetMask("Bounds");
    }
    public void Update(){
        if(Input.GetKey(KeyCode.D)){
            if(!_colliderRight.IsTouchingLayers(mask)){
                _player.MoveTo(Mathf.FloorToInt(transform.position.x+1+0.5f), Mathf.FloorToInt(transform.position.y));
            }
        } else
        if(Input.GetKey(KeyCode.A)){
            if(!_colliderLeft.IsTouchingLayers(mask)){
                _player.MoveTo(Mathf.CeilToInt(transform.position.x-1-0.5f), Mathf.CeilToInt(transform.position.y));
            }
        } else
        if(Input.GetKey(KeyCode.S)){
            if(!_colliderDown.IsTouchingLayers(mask)){
                _player.MoveTo(Mathf.CeilToInt(transform.position.x), Mathf.CeilToInt(transform.position.y-1-0.5f));
            }
        } else
        if(Input.GetKey(KeyCode.W)){
            if(!_colliderUp.IsTouchingLayers(mask)){
                _player.MoveTo(Mathf.FloorToInt(transform.position.x), Mathf.FloorToInt(transform.position.y+1+0.5f));
            }
        }
        
        if(Input.GetKeyDown(KeyCode.F)){

        }
        if(Input.GetKeyDown(KeyCode.G)){
        }
    }
}

