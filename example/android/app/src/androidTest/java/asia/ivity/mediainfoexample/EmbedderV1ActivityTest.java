package asia.ivity.mediainfoexample;

import androidx.test.rule.ActivityTestRule;
import dev.flutter.plugins.e2e.FlutterRunner;
import org.junit.Rule;
import org.junit.runner.RunWith;

@RunWith(FlutterRunner.class)
public class EmbedderV1ActivityTest {
  @Rule
  public ActivityTestRule<EmbedderV1Activity> rule =
      new ActivityTestRule<>(EmbedderV1Activity.class);
}