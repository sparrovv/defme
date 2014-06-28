require "test/unit"
require 'net/http'
require 'json'

class DefMe
  def self.build
    `go build -o test/defme`
  end

  def self.httpd
    @pid = Process.spawn('./test/defme server --port 9879 >> /tmp/defme_test.log 2>&1 &')
  end

  def self.remove_binary
    `rm test/defme`
  end

  def self.kill
    p "kill #{@pid}"
    `kill #{@pid}`
  end
end

class TestDefme < Test::Unit::TestCase
  def setup
    DefMe.build
    DefMe.httpd
  end

  def teardown
    DefMe.kill
    DefMe.remove_binary
  end

  def test_binary_exists
    assert(File.exists?("./test/defme"))
  end

  def test_server_works
    uri = URI('http://localhost:9879')
    uri.query = URI.encode_www_form({ :phrase => "level up", :lang => 'pl' })
    res = Net::HTTP.get_response(uri)

    assert(res.is_a?(Net::HTTPSuccess))

    response = JSON.parse(res.body)
    assert_equal(response['translation'], "poziom w górę")
    assert_equal(response['definitions'], ["To progress to the next level of player character stats and abilities, often by acquiring experience points in role-playing games."])
  end
end
