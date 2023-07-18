import keras
import requests


class Prevue:
    def __init__(
        self,
        user_id: str,
        email: str,
        passwoard: str,
        url: str,
        connector_name: str,
        project_name: str,
        model_name: str,
        architecture: str,
        weights: str,
    
    ) -> None:
        """Define the initial variables for connection.
        """
       
        self.user_id = user_id
        self.email = email
        self.passwoard = passwoard
        self.url = url
        self.connector_name = connector_name
        self.project_name = project_name
        self.model_name = model_name
        self.architecture = architecture
        self.weights = weights
        self.start_session()

    def start_session(self):
        url = f"http://{self.url}/api/connector/session"
        userData = {
            "user_id": self.user_id,
            "email": self.email,
            "passwoard": self.passwoard,
        }

        projectData = {"project_name": self.project_name}
        modelsData = {
            "model_name": self.model_name,
            "connector": self.connector_name,
            "architecture": self.architecture,
            "weights": self.weights,
        }

        data = {
            "userData": userData,
            "modelsData": modelsData,
            "projectData": projectData,
        }

        # get data to the API
        request_post = requests.post(
            url,
            json=data,
        )

    def capture(self, metrics: dict):
        """Capture metrics.

        Connect to backend and send data through prot specified by user.

        Args:
            metrics (dict): Metrics defined in the dict format.
        """

        url = f"http://{self.url}/api/connector/metrics"

        data = {
            "metricsData": metrics,
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
    def __init__(
        self,
        user_id: str,
        email: str,
        passwoard: str,
        url: str,
        connector_name: str,
        project_name: str,
        model_name: str,
        architecture: str,
        weights: str,
    ):
        Prevue.__init__(
            self,
            user_id,
            email,
            passwoard,
            url,
            connector_name,
            project_name,
            model_name,
            architecture,
            weights,
        )

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
