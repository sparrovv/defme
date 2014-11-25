require "minitest/autorun"
require 'net/http'
require 'json'

class DefMe
  def self.build
    `go build -o test/defme`
  end

  def self.port
    return @port if @port

    # find random open port
    server = TCPServer.new('127.0.0.1', 0)
    @port = server.addr[1]
    server.close
    @port
  end

  def self.httpd
    @pid = Process.spawn("./test/defme server --port #{port} >> /tmp/defme_test.log 2>&1 &")
  end

  def self.remove_binary
    `rm test/defme`
  end

  def self.kill
    `kill #{@pid}`
  end
end

class TestDefme < Minitest::Test
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

  def test_server_works_1
    response = _load_response("level up")
    assert(["poziom wyżej", "poziom w górę"].include?(response['translation']))
    assert_equal(
      ["To progress to the next level of player character stats and abilities, often by acquiring experience points in role-playing games."],
      response['definitions']
    )
  end

  def test_server_works_2
    response = _load_response("turnstile")

    assert_equal(
      "bramka",
      response['translation']
    )
    assert_equal(
      ["bramka","kołowrót przy wejściu"],
      response['extraTranslations']
    )
    assert_equal(
      ["A mechanical device used to control passage from one public area to another, typically consisting of several horizontal arms supported by and radially projecting from a central vertical post and allowing only the passage of individuals on foot.","A similar structure that permits the passage of an individual once a charge has been paid or that counts the number of individuals passing through."],
      response['definitions']
    )
    assert_equal(
      nil,
      response['synonyms']
    )
    assert(
      ["They drew three million customers this year, led the league in turnstile clicks, in a place where just a few miles either direction, the main population base is crows.","In real life, however, Afghanistan is as Richard Nixon put it in The Real War, \"has long been a cockpit of great-power intrigue for the same reason that it used to be called the turnstile of Asia's fate\".","Beyond the turnstile was a passage with walls painted white.","When I gestured to my Trasportation Access Pass, he made a rude gesture and waved to the coin turnstile.","Immediately opposite to the turnstile was the open door of a large building; flags surmounted it, and at each side was a large proclamation in red and white."],
      response['examples']
    )
  end

  def _load_response(word)
    uri = URI("http://localhost:#{DefMe.port}")
    uri.query = URI.encode_www_form({:word => word, :to => 'pl'})
    res = Net::HTTP.get_response(uri)

    assert(res.is_a?(Net::HTTPSuccess))

    JSON.parse(res.body)
  end
end
