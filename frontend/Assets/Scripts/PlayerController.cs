using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using Newtonsoft.Json;

public class PlayerController : MonoBehaviour
{
    [SerializeField] PlayerEntity _player;

    int direction = -2;

    [SerializeField] Collider2D _colliderUp;
    [SerializeField] Collider2D _colliderDown;
    [SerializeField] Collider2D _colliderLeft;
    [SerializeField] Collider2D _colliderRight;

    [SerializeField] Collider2D _interactCollider;
    LayerMask boundsMask;
    LayerMask interactMask;
    LayerMask combatMask;
    LayerMask exchangeMask;

    public bool inWorld = true;

    GameObject _bt;

    [SerializeField] WorldCanvasController _worldUI;

    [SerializeField] Bochamon[] bochamones;
    [SerializeField] List<Bochamon> _bochamonsPrefabs;

    public void Start(){
         boundsMask = LayerMask.GetMask("Bounds");
         interactMask = LayerMask.GetMask("Interactable");
         combatMask = LayerMask.GetMask("Combat");
         exchangeMask = LayerMask.GetMask("Exchange");
         _bt = GameObject.FindWithTag("qwe");
         bochamones = new Bochamon[3];

         _player.SetPosition(100,103);
    }

    public void Update(){
        if (inWorld){
            Movement();
        }
    }

    public List<Bochamon> bochamons;
    
    Bochamon b1;
    Bochamon b2;
    Bochamon b3;
    public void SetBochamons(InventoryResponse bochas)
    {
        if (b1 != null){
            Destroy(b1.gameObject);
        }
        if (b2 != null){
            Destroy(b1.gameObject);
        }
        if (b3 != null){
            Destroy(b1.gameObject);
        }
        bochamons.Clear();
        b1 = Instantiate(_bochamonsPrefabs[bochas.value[0]]);
        bochamones[0] = b1;
        bochamons.Add(b1);
        b2 = Instantiate(_bochamonsPrefabs[bochas.value[1]]);
        bochamones[1] = b2;
        bochamons.Add(b2);
        b3 = Instantiate(_bochamonsPrefabs[bochas.value[2]]);
        bochamones[2] = b3;
        bochamons.Add(b3);
    }


    void Movement(){
        if(Input.GetKey(KeyCode.D)){
            if(!_colliderRight.IsTouchingLayers(boundsMask)){
                direction = _player.MoveTo(Mathf.FloorToInt(transform.position.x+1+0.5f), Mathf.FloorToInt(transform.position.y), SendMovement);
            } else {
                direction = _player.GetMovementDirection((int)transform.position.x+1,(int)transform.position.y);
                _player._animator.SetInteger("direction", direction);
                _player._animator.SetBool("moving", false);
            }
            _interactCollider.gameObject.transform.position = new Vector3(transform.position.x + 1,transform.position.y,1);
        } else
        if(Input.GetKey(KeyCode.A)){
            if(!_colliderLeft.IsTouchingLayers(boundsMask)){
                direction = _player.MoveTo(Mathf.CeilToInt(transform.position.x-1-0.5f), Mathf.CeilToInt(transform.position.y), SendMovement);
            }  else {
                direction = _player.GetMovementDirection((int)transform.position.x-1,(int)transform.position.y);
                _player._animator.SetInteger("direction", direction);
                _player._animator.SetBool("moving", false);
            }
            _interactCollider.gameObject.transform.position = new Vector3(transform.position.x - 1,transform.position.y,1);
        } else
        if(Input.GetKey(KeyCode.S)){
            if(!_colliderDown.IsTouchingLayers(boundsMask)){
                direction = _player.MoveTo(Mathf.CeilToInt(transform.position.x), Mathf.CeilToInt(transform.position.y-1-0.5f), SendMovement);
            }  else {
                direction = _player.GetMovementDirection((int)transform.position.x,(int)transform.position.y-1);
                _player._animator.SetInteger("direction", direction);
                _player._animator.SetBool("moving", false);
            }
            _interactCollider.gameObject.transform.position = new Vector3(transform.position.x,transform.position.y-1,1);
        } else
        if(Input.GetKey(KeyCode.W)){
            if(!_colliderUp.IsTouchingLayers(boundsMask)){
                direction = _player.MoveTo(Mathf.FloorToInt(transform.position.x), Mathf.FloorToInt(transform.position.y+1+0.5f), SendMovement);
            }  else {
                direction = _player.GetMovementDirection((int)transform.position.x,(int)transform.position.y+1);
                _player._animator.SetInteger("direction", direction);
                _player._animator.SetBool("moving", false);
            }
            _interactCollider.gameObject.transform.position = new Vector3(transform.position.x,transform.position.y+1,1);
        }


        
        if(Input.GetKeyDown(KeyCode.F)){
            Interact();
        }
        if(Input.GetKeyDown(KeyCode.G)){
        }
    }

    int shouldSendMovement = 0;
    public void SendMovement(int x, int y){
        // if (shouldSendMovement != 0){
        //     shouldSendMovement--;
        //     return;
        // }
        // shouldSendMovement = 3;

        MoveMessage mm = new()
        {
            msgtype = "move",
            x = x,
            y = y
        };
        Connection.Instance.SendWebSocketMessage(JsonConvert.SerializeObject(mm));
    }


    void Interact(){
        if ( _interactCollider.IsTouchingLayers(interactMask)){
            if (bochamons.Count == 0){
                return;
            }
            List<Collider2D> results = new List<Collider2D>();
            ContactFilter2D cf = new ContactFilter2D();
            cf.SetLayerMask(interactMask);
            int amount = _interactCollider.OverlapCollider(cf, results);
            _worldUI.OpenBochamonTrade();
            inWorld = false;      
            return;
        }
        if ( _interactCollider.IsTouchingLayers(exchangeMask)){
            List<Collider2D> results = new List<Collider2D>();
            ContactFilter2D cf = new ContactFilter2D();
            cf.SetLayerMask(exchangeMask);
            int amount = _interactCollider.OverlapCollider(cf, results);
            _worldUI.OpenExchange();
            inWorld = false; 
            return;
        }

        if (_interactCollider.IsTouchingLayers(combatMask)){
            List<Collider2D> results = new List<Collider2D>();
            ContactFilter2D cf = new ContactFilter2D();
            cf.SetLayerMask(combatMask);
            int amount = _interactCollider.OverlapCollider(cf, results);
            PlayerEntity face  = results[0].gameObject.GetComponent<PlayerEntity>();
            Debug.Log(face.uuid);
            DuelRequestMessage duel = new(){
                enemy = face.uuid,
                msgtype = "duelrequest"
            };
            Connection.Instance.SendWebSocketMessage(JsonConvert.SerializeObject(duel));
            return;
        }
    }
}

