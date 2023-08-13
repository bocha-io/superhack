using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.SceneManagement;
using Newtonsoft.Json;

public class Gamemaster : MonoBehaviour
{
    [SerializeField] PlayerEntity _playerPrefab;
    [SerializeField] PlayerController _mainPlayer;
    Dictionary<string, PlayerEntity> _players = new Dictionary<string, PlayerEntity>();

    [SerializeField] WorldCanvasController _worldUI;

    (string, string) message;
    public void Update()
    {
        if(Connection.Instance.messages.TryDequeue(out message)){
            ExecuteMessage(message);
        }
    }

    public void Start(){
        PublicBattleStatus.valid = false;
        StartCoroutine(SendInventory());
    }

    public IEnumerator SendInventory(){
        yield return new WaitForSeconds(5);
        BaseMsg msg = new(){
            msgtype = "inventory"
        };
        Connection.Instance.SendWebSocketMessage(JsonConvert.SerializeObject(msg));
    }

    public void ExecuteMessage((string type, string content) message)
    {
        switch(message.type)
        {
            case "duelrequestresponse":
            {
                DuelRequestMessageResponse duelRequestResponse = JsonConvert.DeserializeObject<DuelRequestMessageResponse>(message.content);
                if (duelRequestResponse.value.playera == Connection.Instance.wallet){
                    // Show Waiting for response
                    Debug.Log("Waiting For Message");
                } else {
                    if (duelRequestResponse.value.playerb == Connection.Instance.wallet){
                        _worldUI.ChallengedToDuel(duelRequestResponse.value.playera);
                    }
                }
                break;
            }
            case "duelresponseresponse":            {
                DuelResponseMessageResponse response = JsonConvert.DeserializeObject<DuelResponseMessageResponse>(message.content);
                if (response.value.playera == Connection.Instance.wallet || response.value.playerb == Connection.Instance.wallet){
                    SceneManager.LoadScene("BattleScene");
                }
                break;
            }
            case "battlestatus":
            {
                BattleStatus battleStatus = JsonConvert.DeserializeObject<BattleStatus>(message.content);
                PublicBattleStatus.status = battleStatus;
                PublicBattleStatus.valid = true;
                break;
             
            }
            case "swapresponse":
            {
                SwapMessageResponse swapResponse = JsonConvert.DeserializeObject<SwapMessageResponse>(message.content);
                // UI canje realizado
                BaseMsg msg = new(){
                    msgtype = "inventory"
                };
                Connection.Instance.SendWebSocketMessage(JsonConvert.SerializeObject(msg));
                break;
            }
            case "inventoryresponse":
            {
                InventoryResponse inventoryResponse = JsonConvert.DeserializeObject<InventoryResponse>(message.content);
                
                if (inventoryResponse.value != null && inventoryResponse.value.Length!=0){
                    // UI canje realizado
                    _mainPlayer.SetBochamons(inventoryResponse);
                }
                else {
                    // ask again in 5
                    StartCoroutine(SendInventory());
                }
                break;
            }
            case "mapstatus":
            {
                MapStatus mapStatus = JsonConvert.DeserializeObject<MapStatus>(message.content);
                PlaceInWorld(mapStatus);
                break;
            }
            case "bridgeresponse":
            {
                BridgeMessageResponse bridgeMessageResponse = JsonConvert.DeserializeObject<BridgeMessageResponse>(message.content);
                // if (bridgeMessageResponse.value != null && bridgeMessageResponse.value != ""){
                //     Application.OpenURL("https://optimistic.etherscan.io/tx/"+bridgeMessageResponse.value);
                // }
                break;
            }
            case "move":
            {
                //
                break;
            }
            case "moveresponse":
            {
                // Verify movement
                break;
            }
            case "creatematchresponse":
            {
                // Load Battle Scene
                SceneManager.LoadScene("BattleScene");
                break;
            }
            default:
            {
                Debug.LogWarning("Message not recognized: " + message.type);
                break;
            }
        }
    }

    public void PlaceInWorld(MapStatus mapStatus){
        Debug.Log(mapStatus.playerspos);
        // TODO: Handle Disconnected players
        foreach(PlayerPos p in mapStatus.playerspos){
            string temp = p.playerid.TrimStart('0');
            temp = temp.TrimStart('x');
            temp = temp.TrimStart('0');
            temp = "0x" + temp;
            if (temp != Connection.Instance.wallet)
            {
                if (! _players.ContainsKey(temp)) {
                    PlayerEntity player = Instantiate(_playerPrefab, transform);
                    player.uuid = temp;
                    player.SetPosition(p.X, p.Y);
                    _players.Add(temp, player);
                } else {
                    _players[temp].MoveTo(p.X, p.Y, (int a, int b) => { });
                }
            }
        }
    }
}
