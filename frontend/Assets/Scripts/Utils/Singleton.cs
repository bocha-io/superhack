using UnityEngine;

public class Singleton<T> : MonoBehaviour where T : Component
{
    private static T _instance;
    public static T Instance
    {
        get
        {
            if (_instance == null)
            {
                var objs = FindObjectOfType(typeof(T)) as T[];
                if (objs != null)
                {
                    if (objs.Length > 0)
                    {
                        _instance = objs[0];
                    }
                    if (objs.Length > 1)
                    {
                        Debug.Log("There is more than one " + typeof(T).Name);

                    }
                }
                if (_instance == null)
                {
                    GameObject obj = new GameObject();
                    obj.hideFlags = HideFlags.HideAndDontSave;
                    _instance = obj.AddComponent<T>();
                }
            }
            return _instance;
        }
    }

    public virtual void Awake()
    {
        // create the instance
        if (_instance == null)
        {
            _instance = this as T;
            DontDestroyOnLoad(this.gameObject);
        }
        else
        {
            Destroy(gameObject);
        }
    }
}