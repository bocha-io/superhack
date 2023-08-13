using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.SceneManagement;
using TMPro;
using Newtonsoft.Json;

public class LoginController : MonoBehaviour
{
    [SerializeField] TMP_InputField _username;
    [SerializeField] TMP_InputField _password;

    public void Login(){
        Connection.Instance.Connect(_username.text, _password.text);
    }

    (string, string) message;
    public void Update(){
        if(Connection.Instance.messages.TryDequeue(out message)){
            ExecuteMessage(message);
        }
    }

    public void ExecuteMessage((string type, string content) message){
        switch(message.type){
            case "connectresponse":
            {
                ConnectResponse connect = JsonConvert.DeserializeObject<ConnectResponse>(message.content);
                Connection.Instance.wallet = connect.value;
                SceneManager.LoadScene("WorldScene");
                break;
            }
            default:
            {
                Debug.LogWarning("Unrecognized Message");
                break;
            }
        }
    }
}
