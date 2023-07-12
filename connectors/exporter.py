import keras
import requests


class Prevue:
    def __init__(
        self,
        connector_name: str,
        project_name: str,
        user_id: str,
        url: str,
        **kwargs,
    ) -> None:
        """Define the initial variables for connection.

        Args:
            connector_name (str): Connector name.
            project_name (str): Project name.
            uid (str): User id.
            url (str): Url.
        """
        super().__init__(**kwargs)
        self.connector_name = connector_name
        self.project_name = project_name
        self.user_id = user_id
        self.url = url
        print("test")

        # NOTE add user name or authenticator

    def capture(self, metrics: dict):
        """Capture metrics.

        Connect to backend and send data through prot specified by user.

        Args:
            metrics (dict): Metrics defined in the dict format.
        """

        url = f"http://{self.url}/api/connector/metrics"
        userData = {
            "user_id": self.user_id,
            "email": "kubinsmarta77@gmail.com",
            "passwoard": "marta1",
        }

        projectData = {"project_name": self.project_name}
        modelsData = {
            "model_name": "marta1",
            "connector": self.connector_name,
            "architecture": "architecture test",
            "weights": "weights test",
        }

        data = {
            "metricsData": metrics,
            "userData": userData,
            "modelsData": modelsData,
            "projectData": projectData,
        }

        # print(data)

        # get data to the API
        request_post = requests.post(
            url,
            json=data,
        )

        # print(request_post.json())
        return
        

class PrevueKerasCallback(Prevue, keras.callbacks.Callback):
    def __init__(self, connector_name: str, project_name: str, user_id: str, url: str):
        Prevue.__init__(self, connector_name, project_name, user_id, url)

    # def on_train_batch_end(self, batch, logs=None):
    #     self.capture({"batch": batch, "loss": logs["loss"]})

    # def on_test_batch_end(self, batch, logs=None):
    #     self.capture({"batch": batch, "loss": logs["loss"]})

    def on_epoch_end(self, epoch, logs=None):
        self.capture(
            {
                "epoch": epoch,
                "batch": 1,
                "loss_name": "loss",
                "loss_value": logs["loss"],
                "metrics_name": "accuracy",
                "metrics_value": logs["accuracy"],
            }
        )
