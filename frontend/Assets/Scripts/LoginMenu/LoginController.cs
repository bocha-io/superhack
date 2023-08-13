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

    [SerializeField] GameObject _waitForLogin;
    [SerializeField] GameObject _login;

    public void Login(){
        Connection.Instance.Connect(_username.text, _password.text);
    }

    public IEnumerator WaitToLogin(){
        _login.SetActive(false);
        _waitForLogin.SetActive(true);
        yield return new WaitForSeconds(10);
        SceneManager.LoadScene("WorldScene");
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
                StartCoroutine(WaitToLogin());
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
