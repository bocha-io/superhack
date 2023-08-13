using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class LinkExplorerAction : MenuOptionAction
{
    public override void Execute()
    {
        Application.OpenURL("https://optimistic.etherscan.io/address/"+Connection.Instance.wallet);
    }
}
