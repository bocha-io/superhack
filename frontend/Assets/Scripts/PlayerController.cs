using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using Newtonsoft.Json;

public class PlayerController : MonoBehaviour
{
    [SerializeField] PlayerEntity _player;
    [SerializeField] Collider2D _colliderUp;
    [SerializeField] Collider2D _colliderDown;
    [SerializeField] Collider2D _colliderLeft;
    [SerializeField] Collider2D _colliderRight;
    LayerMask boundsMask;
    LayerMask interactMask;
    public void Start(){
         boundsMask = LayerMask.GetMask("Bounds");
         interactMask = LayerMask.GetMask("Interactable");
    }

    public void Update(){
        Movement();
    }

    void Movement(){
        if(Input.GetKey(KeyCode.D)){
            if(!_colliderRight.IsTouchingLayers(boundsMask)){
                _player.MoveTo(Mathf.FloorToInt(transform.position.x+1+0.5f), Mathf.FloorToInt(transform.position.y), SendMovement);
            }
        } else
        if(Input.GetKey(KeyCode.A)){
            if(!_colliderLeft.IsTouchingLayers(boundsMask)){
                _player.MoveTo(Mathf.CeilToInt(transform.position.x-1-0.5f), Mathf.CeilToInt(transform.position.y), SendMovement);
            }
        } else
        if(Input.GetKey(KeyCode.S)){
            if(!_colliderDown.IsTouchingLayers(boundsMask)){
                _player.MoveTo(Mathf.CeilToInt(transform.position.x), Mathf.CeilToInt(transform.position.y-1-0.5f), SendMovement);
            }
        } else
        if(Input.GetKey(KeyCode.W)){
            if(!_colliderUp.IsTouchingLayers(boundsMask)){
                _player.MoveTo(Mathf.FloorToInt(transform.position.x), Mathf.FloorToInt(transform.position.y+1+0.5f), SendMovement);
            }
        }
        
        if(Input.GetKeyDown(KeyCode.F)){
            Interact();
        }
        if(Input.GetKeyDown(KeyCode.G)){
        }
    }

    void SendMovement(int x, int y){
        MoveMessage mm = new()
        {
            msgtype = "move",
            x = x,
            y = y
        };
        // Connection.Instance.SendWebSocketMessage(JsonConvert.SerializeObject(mm));
    }

    void Interact(){
        if ( _colliderUp.IsTouchingLayers(boundsMask) ||
        _colliderDown.IsTouchingLayers(boundsMask) ||
        _colliderLeft.IsTouchingLayers(boundsMask) ||
        _colliderRight.IsTouchingLayers(boundsMask) 
        ){
            Debug.Log("Interact");
        }
    }
}

