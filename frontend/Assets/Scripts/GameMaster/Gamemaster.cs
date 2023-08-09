using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using Newtonsoft.Json;

public class Gamemaster : MonoBehaviour
{
    [SerializeField] PlayerEntity _playerPrefab;
    [SerializeField] PlayerController _mainPlayer;
    Dictionary<string, PlayerEntity> _players;

    (string, string) message;
    public void Update()
    {
        // if(Connection.Instance.messages.TryDequeue(out message)){
        //     ExecuteMessage(message);
        // }
    }

    public void ExecuteMessage((string type, string content) message)
    {
        switch(message.type)
        {
            case "mapstatus":
            {
                MapStatus mapStatus = JsonConvert.DeserializeObject<MapStatus>(message.content);
                InstantiateWorld(mapStatus);
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
            default:
            {
                Debug.LogWarning("Message not recognized: " + message.type);
                break;
            }
        }
    }

    public void InstantiateWorld(MapStatus mapStatus){
        foreach(PlayerPos p in mapStatus.playerpos){
            if (! _players.ContainsKey(p.id)) {
                PlayerEntity player = Instantiate(_playerPrefab, transform);
                //  player.SetPosition(p.x, p.y);
                _players.Add(p.id, player);
            } else {
                //  _players[p.id].SetPosition(p.x, p.y);
            }
            
        }
    }
}
