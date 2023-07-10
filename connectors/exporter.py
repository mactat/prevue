import keras
import requests


class Prevue:
    def __init__(
        self, connector_name: str, project_name: str, uid: str, url: str, **kwargs,
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
        self.uid = uid
        self.url = url
        print('test')

        # NOTE add user name or authenticator

    def capture(self, metrics: dict):
        """Capture metrics.

        Connect to backend and send data through prot specified by user.

        Args:
            metrics (dict): Metrics defined in the dict format.
        """

        url = f"http://{self.url}/api/connector/metrics"

        data = {
            "connectorName": self.connector_name,
            "metricsData": metrics,
            "projectName": self.project_name,
            "uid": self.uid,
        }

        # get data to the API
        post_response = requests.post(
            url,
            json=data,
        )

        return

class PrevueKerasCallback(Prevue, keras.callbacks.Callback):
    def __init__(self, connector_name: str, project_name: str, uid: str, url: str):
        Prevue.__init__(self, connector_name, project_name, uid, url)

    # def on_train_batch_end(self, batch, logs=None):
    #     self.capture({"batch": batch, "loss": logs["loss"]})

    # def on_test_batch_end(self, batch, logs=None):
    #     self.capture({"batch": batch, "loss": logs["loss"]})

    # def on_epoch_end(self, epoch, logs=None):
    #     self.capture({"epoch": epoch, "loss": logs})

    def on_epoch_end(self, epoch, logs=None):
        self.capture({"accuracy": logs["accuracy"], "mse": logs["loss"]})
