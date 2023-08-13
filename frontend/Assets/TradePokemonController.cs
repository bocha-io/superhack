using System.Collections;
using System.Collections.Generic;
using UnityEngine.SceneManagement;
using UnityEngine;
using Newtonsoft.Json;

public class TradePokemonController : OptionsController
{
    public List<MenuOption> _allOptions;
    int _tradeBochamon;

    [SerializeField] BochamonPanelController _previousPanel;

    public void Setup(List<Bochamon> bochamons, int tradeBochamon)
    {
        _tradeBochamon = tradeBochamon;
        options.Clear();
        foreach (MenuOption option in _allOptions){
            bool discarded = false;
            foreach (Bochamon b in bochamons){
                if (option.GetComponent<ExchangeBochamonAction>().montype == b.montype){
                    option.gameObject.SetActive(false);
                    discarded = true;
                }
            }
            if (!discarded){
                options.Add(option);
                option.gameObject.SetActive(true);
            }
                
        }
    }
    

    public override void GoBack()
    {
        _previousPanel.enable = true;
        this.gameObject.SetActive(false);
    }

    public void PerformSwap(int montype){
        Debug.Log("Trade " + _tradeBochamon + " for " + montype);
        SwapMessage swap = new(){
            msgtype="swap",
            montype= montype,
            pos= _tradeBochamon
        };
        Connection.Instance.SendWebSocketMessage(JsonConvert.SerializeObject(swap));
    }

}
