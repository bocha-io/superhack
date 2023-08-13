using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class Player : MonoBehaviour
{
    public List<Bochamon> bochamons;
    public string id;
    

    public Bochamon GetBochamon(string id){
        foreach (Bochamon b in bochamons){
            if (b.uuid == id){
                return b;
            }
        }
        return null;
    }
}
